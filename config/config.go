package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/antigloss/go/logger"
)

var (
	Config *config
)

func init() {
	getConfig()
}

type config struct {
	Server_Port string `json:"server_port"`
	GRPC_Port   string `json:"grpc_port"`
}

func getConfig() *config {

	if Config == nil {
		data, err := ioutil.ReadFile("grpc_server_demo.json")
		if err != nil {
			logger.Error("Failed to parse config file %s" + err.Error())
		}
		err = json.Unmarshal(data, &Config)
		if err != nil {
			logger.Error("Failed to unmarshal config file %s" + err.Error())
		}
	}
	return Config
}
