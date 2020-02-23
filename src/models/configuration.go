package models

import (
	"encoding/json"
	"os"
	"path/filepath"
)

var (
	//Configuration defines config variables
	Configuration Config
)

// Config defines configuration variables
type Config struct {
	//Db denotes database configuration
	Db struct {
		EndPoint string `json:"endpoint"`
	} `json:"db"`
	//Server denotes database configuration
	Server struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"server"`
}

//New creates a new instance of Config
func New(fileName string) Config {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	config := Config{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)

	if err != nil {
		panic(err)
	}

	return config
}

//init sets up the initial states
func init() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	Configuration = New(filepath.Join(dir, "config", "app.config.json"))
}
