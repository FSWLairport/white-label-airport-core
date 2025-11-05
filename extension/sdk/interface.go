package sdk

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"

	"github.com/pppwaw/white-label-airport-core/config"
	v2 "github.com/pppwaw/white-label-airport-core/v2"
	"github.com/sagernet/sing-box/option"
)

func RunInstance(whitelabelairportSettings *config.WhiteLabelAirportOptions, singconfig *option.Options) (*v2.WhiteLabelAirportService, error) {
	return v2.RunInstance(whitelabelairportSettings, singconfig)
}

func ParseConfig(whitelabelairportSettings *config.WhiteLabelAirportOptions, configStr string) (*option.Options, error) {
	if whitelabelairportSettings == nil {
		whitelabelairportSettings = config.DefaultWhiteLabelAirportOptions()
	}
	if strings.HasPrefix(configStr, "http://") || strings.HasPrefix(configStr, "https://") {
		client := &http.Client{}
		configPath := strings.Split(configStr, "\n")[0]
		// Create a new request
		req, err := http.NewRequest("GET", configPath, nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			return nil, err
		}
		req.Header.Set("User-Agent", "WhiteLabelAirport/2.3.1 ("+runtime.GOOS+") like ClashMeta v2ray sing-box")
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making GET request:", err)
			return nil, err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read config body: %w", err)
		}
		configStr = string(body)
	}
	return config.ParseConfigContentToOptions(configStr, true, whitelabelairportSettings, false)
}
