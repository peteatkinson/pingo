package core

import (
	"encoding/json"
	"io/ioutil"
)

func LoadServiceConfig() []ServiceConfig {
	var serviceConfigs []ServiceConfig

	data, err := ioutil.ReadFile("./schedule.json")

	if err != nil {
		return []ServiceConfig{}
	}

	err = json.Unmarshal(data, &serviceConfigs)

	if err != nil {
		return []ServiceConfig{}
	}

	return serviceConfigs
}
