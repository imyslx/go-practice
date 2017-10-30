package main

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

	zlog "github.com/rs/zerolog/log"
)

type Settings struct {
	Basic BasicSettings
	Items []MonitoringItem
}

type BasicSettings struct {
	ServerName string
	ClientName string
}

type MonitoringItem struct {
	Key      string
	Command  string
	Duration int
	DataType string
}

func GetLocalConfig(file string) (Settings, error) {

	data := new(Settings)

	// Read config file.
	//     ([]bytes, err)
	jsonBytes, err := ioutil.ReadFile("test.json")
	if err != nil {
		return *data, err
	}

	if err := json.Unmarshal(jsonBytes, data); err != nil {
		return *data, err
	}

	return *data, nil
}

func ValidateBasicSettings(config BasicSettings) bool {
	errFlag := 0

	if config.ServerName == "" {
		zlog.Error().Msg("ServerName config is Empty. Please define ServerName.")
		errFlag = 1
	}
	if config.ClientName == "" {
		zlog.Error().Msg("ClientName config is Empty. Please define ClientName.")
		errFlag = 1
	}

	if errFlag == 1 {
		zlog.Fatal().Msg("Detected errors in Basic config.")
		return false
	}

	return true
}

func ValidateMonitoringItem(items []MonitoringItem) bool {
	errFlag := 0
	for i, item := range items {
		itemBytes, err := json.Marshal(item)
		if err != nil {
			zlog.Fatal().Err(err).Msg("JSON Marshal error.")
		}

		if item.Key == "" {
			zlog.Error().
				Str("Data", string(itemBytes)).
				Msg("Key is Empty in MonitoringItem [" + strconv.Itoa(i) + "].")
			errFlag = 1
		}
		if item.Command == "" {
			zlog.Error().
				Str("Data", string(itemBytes)).
				Msg("key " + item.Key + " : Command is Empty.")
			errFlag = 1
		}
		if item.Duration == 0 {
			zlog.Info().
				Str("Data", string(itemBytes)).
				Msg("key " + item.Key + " : Duration is Empty.\nSet to default 60.")
			item.Duration = 60
		}
		if item.DataType == "" {
			zlog.Error().
				Str("Data", string(itemBytes)).
				Msg("key " + item.Key + " : DataType is Empty.")
			errFlag = 1
		}
	}

	if errFlag == 1 {
		zlog.Error().Msg("Detected errors in Monitoring Items.")
		return false
	}

	return true
}

func Configuration(config Settings) {
	zlog.Debug().Msg("Starting config validation.")

	vbs := ValidateBasicSettings(config.Basic)

	//NOT IMPLEMENTED: Sync MonitoringItem to server.
	//sc := SyncConfig()

	vmi := ValidateMonitoringItem(config.Items)

	zlog.Debug().Msg("Completed config validation.")
	if vbs == false || vmi == false {
		zlog.Fatal().Msg("Some Config Has error. Check it !!")
	}

	zlog.Debug().Msg("No problem in all configs !!")
}

func main() {
	zlog.Info().Msg("Start to process...")

	// Get local config.
	// Then request to sync monitoring items to server. (not implemented)
	file := "test.json"
	config, err := GetLocalConfig(file)
	if err != nil {
		zlog.Fatal().Err(err).Msg("Could not get config. ")
	}
	zlog.Info().Msg("Successful to get config. Continue.")

	// Sync and Validate config.
	Configuration(config)

}
