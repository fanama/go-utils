package database

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Port     int    `json:"port"`
	DB       string `json:"db"`
	Host     string `json:"host"`
}

func CreateConfig(fileName string) (conf Configuration, err error) {
	fmt.Println("CreatingFile...")

	file, err := os.Create(fileName + ".json")
	if err != nil {
		return conf, err
	}

	defer file.Close()
	fmt.Print("Enter Host:")
	fmt.Scan(&conf.Host)
	fmt.Print("Enter Username: ")
	fmt.Scan(&conf.User)
	fmt.Print("Enter Password: ")
	fmt.Scan(&conf.Password)
	fmt.Print("Enter  port : ")
	fmt.Scan(&conf.Port)
	fmt.Print("Enter Database: ")
	fmt.Scan(&conf.DB)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "	")
	err = encoder.Encode(conf)
	if err != nil {
		return conf, err
	}

	return conf, err
}

func ReadConfig(filePath string) (conf Configuration, err error) {
	fmt.Println("ReadingFiles...")

	file, err := os.Open(filePath + ".json")

	if err != nil {
		return CreateConfig(filePath)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&conf)
	if err != nil {
		return conf, err
	}

	return conf, err

}
