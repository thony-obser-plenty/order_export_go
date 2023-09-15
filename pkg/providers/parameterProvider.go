package providers

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Parameters struct {
	ApiBaseUri    string `yaml:"api_base_uri"`
	ApiBaseScheme string `yaml:"api_base_scheme"`
	ApiKey        string `yaml:"api_key"`
}

func GetParameters(configPath string) (*Parameters, error) {
	var parameters Parameters

	yamlFile, err := ioutil.ReadFile(configPath)

	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &parameters)

	return &parameters, nil
}
