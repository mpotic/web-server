package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Database struct {
		ConnectionString string `json:"connectionString"`
	} `json:"database"`
}

func getConfigFilePath() string {
	configFilePath, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting the configuration file path! ", err)

	}
	configFilePath += "/config.json"

	return configFilePath
}

func GetConfig() *Config {
	configFilePath := getConfigFilePath()
	configFile, err := os.Open(configFilePath)
	if err != nil {
		log.Fatal("Error opening configuration file!", err)
		return nil
	}

	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&configFile)
	if err != nil {
		log.Fatal("Error decoding the .json file!", err)
		return nil
	}

	var config *Config = new(Config)
	err = decoder.Decode(config)

	return config
}
