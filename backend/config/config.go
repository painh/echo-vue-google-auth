package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var Config map[string]interface{} = ReadConfig()

func ReadConfig() map[string]interface{} {

	// read in YAML file
	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}

	// create map for YAML file
	var yamlMap map[string]interface{}

	// unmarshal YAML file into map

	err = yaml.Unmarshal(yamlFile, &yamlMap)

	if err != nil {

		panic(err)
	}

	return yamlMap
}
