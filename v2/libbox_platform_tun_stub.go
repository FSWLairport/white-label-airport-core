//go:build (!linux && !darwin) && !android && !ios

package v2

import (
	"github.com/sagernet/sing-box/experimental/libbox"
	"github.com/sagernet/sing/common/control"
)

func openTunDevice(_ libbox.TunOptions, _ control.InterfaceFinder) (int32, error) {
	return 0, errUnsupportedPlatformFeature
}
