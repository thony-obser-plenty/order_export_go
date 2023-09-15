package providers

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Routes struct {
	ApiAuthenticatorRoute Route `yaml:"api_authenticator"`
	ApiOrderRoute         Route `yaml:"api_orders"`
}

type Route struct {
	Path   string
	Method string
}

func GetRoutes(configPath string) (*Routes, error) {
	var Routes Routes

	yamlFile, err := ioutil.ReadFile(configPath)

	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &Routes)

	return &Routes, nil
}
