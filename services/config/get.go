package config

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
	DBType     string `json:dbType`
	DBHost     string `json:dbHost`
	DBPort     string `json:dbPort`
	DBName     string `json:dbName`
	DBUser     string `json:dbUser`
	DBPassword string `json:dbPassword`
	SigningKey string `json:signingKey`
}

var config *ConfigProvider
var configOnce sync.Once

func Get() *ConfigProvider {
	configOnce.Do(func() {
		jsonFile, err := os.Open("static/config.json")
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
