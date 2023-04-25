package config_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"budget-tracker/config"
	"budget-tracker/models"
)

type ConfigTestSuite struct {
	suite.Suite
}


func (suite *ConfigTestSuite) TestReadConfig(){

	config, err := config.ReadConfig("test_file/config.yml")
	suite.Require().NoError(err, "no config.yml file found")

	expectedConfig := models.Config{
        File: struct {
            Path     string `yaml:"path"`
        }{
            Path: "test_path",
        },
    }

	suite.Nil(err)
	suite.Equal(config, expectedConfig)
}


func (suite *ConfigTestSuite) TestReadConfig_NoFile(){

	config, err := config.ReadConfig("wrong_path_to_file")
	suite.Require().Error(err)

	suite.Equal(config, models.Config{})
	suite.Error(err, "no config.yml file found")
}

func TestReadConfigSuite(t *testing.T) {
    suite.Run(t, new(ConfigTestSuite))
}
