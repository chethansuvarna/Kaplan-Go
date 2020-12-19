package config

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

type Reader interface {
	ReadConfig() (map[string]interface{}, error)
}

type configReader struct {
	fileName string
}

func NewConfigReader(fileName string) Reader {
	return configReader{fileName: fileName}
}

func (cl configReader) ReadConfig() (map[string]interface{}, error) {
	config := make(map[string]interface{})
	file, err := os.Open(cl.fileName)
	defer file.Close()
	if err != nil {
		log.Println("error while opening config file", err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Println("error while decoding config", err)
		return config, errors.New("error while decoding config")
	}
	return config, nil
}
