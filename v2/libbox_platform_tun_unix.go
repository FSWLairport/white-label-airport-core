//go:build (linux || darwin) && !android && !ios

package v2

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"

	"github.com/sagernet/sing-box/experimental/libbox"
	tun "github.com/sagernet/sing-tun"
	"github.com/sagernet/sing/common/control"

	"golang.org/x/sys/unix"
)

func openTunDevice(options libbox.TunOptions, finder control.InterfaceFinder) (int32, error) {
	rawOptions, err := cloneTunOptions(options, finder)
	if err != nil {
		return 0, err
	}
	tunInterface, err := tun.New(*rawOptions)
	if err != nil {
		return 0, err
	}
	if err = tunInterface.Start(); err != nil {
		_ = tunInterface.Close()
		return 0, err
	}
	fd, err := tunDescriptorFromInstance(tunInterface)
	if err != nil {
		_ = tunInterface.Close()
		return 0, err
	}
	dupFd, err := unix.Dup(fd)
	if err != nil {
		_ = tunInterface.Close()
		return 0, fmt.Errorf("duplicate tun fd: %w", err)
	}
	suppressTunCleanup(tunInterface, fd)
	if err = tunInterface.Close(); err != nil {
		return 0, fmt.Errorf("close tun handle: %w", err)
	}
	return int32(dupFd), nil
}

func cloneTunOptions(options libbox.TunOptions, finder control.InterfaceFinder) (*tun.Options, error) {
	if options == nil {
		return nil, fmt.Errorf("nil tun options")
	}
	value := reflect.ValueOf(options)
	for value.Kind() == reflect.Interface || value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return nil, fmt.Errorf("invalid tun options")
		}
		value = value.Elem()
	}
	field := value.FieldByName("Options")
	if !field.IsValid() || field.IsNil() {
		return nil, fmt.Errorf("tun options missing raw Options")
	}
	rawOptions, ok := field.Interface().(*tun.Options)
	if !ok || rawOptions == nil {
		return nil, fmt.Errorf("invalid tun options payload")
	}
	cloned := *rawOptions
	cloned.FileDescriptor = 0
	if cloned.Name == "" {
		cloned.Name = tun.CalculateInterfaceName("")
	}
	if cloned.InterfaceFinder == nil {
		cloned.InterfaceFinder = finder
	}
	return &cloned, nil
}

func tunDescriptorFromInstance(tunInterface tun.Tun) (int, error) {
	value := reflect.ValueOf(tunInterface)
	if value.Kind() != reflect.Pointer || value.IsNil() {
		return 0, fmt.Errorf("unexpected tun implementation %T", tunInterface)
	}
	elem := value.Elem()
	fdField := elem.FieldByName("tunFd")
	if !fdField.IsValid() {
		return 0, fmt.Errorf("tun implementation %T is missing tunFd", tunInterface)
	}
	fd := int(reflect.NewAt(fdField.Type(), unsafe.Pointer(fdField.UnsafeAddr())).Elem().Int())
	return fd, nil
}

func suppressTunCleanup(tunInterface tun.Tun, fd int) {
	value := reflect.ValueOf(tunInterface)
	if value.Kind() != reflect.Pointer || value.IsNil() {
		return
	}
	elem := value.Elem()
	switch runtime.GOOS {
	case "linux":
		optionsField := elem.FieldByName("options")
		if !optionsField.IsValid() {
			return
		}
		optionValue := reflect.NewAt(optionsField.Type(), unsafe.Pointer(optionsField.UnsafeAddr())).Elem()
		fdField := optionValue.FieldByName("FileDescriptor")
		if fdField.IsValid() && fdField.CanSet() {
			fdField.SetInt(int64(fd))
		}
	case "darwin":
		routeSetField := elem.FieldByName("routeSet")
		if routeSetField.IsValid() && routeSetField.CanSet() {
			routeSetField.SetBool(false)
		}
	}
}
