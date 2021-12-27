package mail

import (
	"encoding/json"
	"fmt"
	"os"
)

func CreateConfig(fileName string) (conf Configuration, err error) {
	fmt.Println("CreatingFile...")

	file, err := os.Create(fileName + ".json")
	if err != nil {
		return conf, err
	}

	defer file.Close()
	fmt.Print("Enter your mail address :")
	fmt.Scan(&conf.Address)
	fmt.Print("Enter your password :")
	fmt.Scan(&conf.Password)

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
