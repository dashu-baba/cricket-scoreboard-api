package models

import (
	"os"
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

	config := Config{}
	config.Db.EndPoint = os.Getenv("DB_CONN")
	config.Server.Host = os.Getenv("HOST")
	config.Server.Port = os.Getenv("PORT")

	return config
}
