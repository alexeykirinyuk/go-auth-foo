package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const configFileName = "config.json"

type Config struct {
	Port             int
	ConnectionString string
}

func GetConfig() (c Config, err error) {
	content, err := os.Open(configFileName)
	if err != nil {
		err = fmt.Errorf("can't open config file: %s", err)
		return
	}

	defer func() {
		if err := content.Close(); err != nil {
			err = fmt.Errorf("can't close config file: %s", err)
			panic(err)
		}
	}()

	byteArray, err := ioutil.ReadAll(content)
	if err != nil {
		err = fmt.Errorf("can't read byte array from config file: %s", err)
		return
	}

	err = json.Unmarshal(byteArray, &c)
	if err != nil {
		err = fmt.Errorf("can't read byte array from config file: %s", err)
		return
	}

	return
}
