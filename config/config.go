package config

import (
	"budget-tracker/models"
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)


func ReadConfig(configFile string) (models.Config, error) {
    configData, err := os.ReadFile(configFile)
    if err != nil {
        return models.Config{}, errors.New("no config.yml file found")
    }
    
    var config models.Config
    err = yaml.Unmarshal(configData, &config)
    if err != nil {
        return models.Config{}, errors.New("not possible to unmarshal config")
    }
    return config, nil
}