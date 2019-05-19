package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

type ConfigProvider struct {
	APIName    string `json:apiName`
	APIVersion string `json:apiVersion`
	HTTPPort   string `json:httpPort`
}

var config *ConfigProvider
var configOnce sync.Once

func GetConfig() *ConfigProvider {
	configOnce.Do(func() {
		jsonFile, err := os.Open("src/go-api-template/config.json")
		if err != nil {
			log.Print("Json open error: ", err)
		}

		byteValue, _ := ioutil.ReadAll(jsonFile)
		if err != nil {
			log.Print("Json read error: ", err)
		}

		json.Unmarshal(byteValue, &config)
	})

	return config
}
