package mobile

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/pppwaw/white-label-airport-core/config"

	v2 "github.com/pppwaw/white-label-airport-core/v2"

	_ "github.com/sagernet/gomobile"
	"github.com/sagernet/sing-box/option"
)

func Setup(baseDir string, workingDir string, tempDir string, debug bool) error {
	return v2.Setup(baseDir, workingDir, tempDir, 0, debug)
	// return v2.Start(17078)
}

func Parse(path string, tempPath string, debug bool) error {
	config, err := config.ParseConfig(tempPath, debug)
	if err != nil {
		return err
	}
	return os.WriteFile(path, config, 0o644)
}

func BuildConfig(path string, WhiteLabelAirportOptionsJson string) (string, error) {
	os.Chdir(filepath.Dir(path))
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	var options option.Options
	err = options.UnmarshalJSON(fileContent)
	if err != nil {
		return "", err
	}
	WhiteLabelAirportOptions := &config.WhiteLabelAirportOptions{}
	err = json.Unmarshal([]byte(WhiteLabelAirportOptionsJson), WhiteLabelAirportOptions)
	if err != nil {
		return "", nil
	}
	if WhiteLabelAirportOptions.Warp.WireguardConfigStr != "" {
		err := json.Unmarshal([]byte(WhiteLabelAirportOptions.Warp.WireguardConfigStr), &WhiteLabelAirportOptions.Warp.WireguardConfig)
		if err != nil {
			return "", err
		}
	}

	if WhiteLabelAirportOptions.Warp2.WireguardConfigStr != "" {
		err := json.Unmarshal([]byte(WhiteLabelAirportOptions.Warp2.WireguardConfigStr), &WhiteLabelAirportOptions.Warp2.WireguardConfig)
		if err != nil {
			return "", err
		}
	}

	return config.BuildConfigJson(*WhiteLabelAirportOptions, options)
}

func GenerateWarpConfig(licenseKey string, accountId string, accessToken string) (string, error) {
	return config.GenerateWarpAccount(licenseKey, accountId, accessToken)
}
