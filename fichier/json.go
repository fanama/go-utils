package fichier

import (
	"encoding/json"
	"fmt"
	"os"
)

func CreateJSON(fileName string) (conf map[string]string, err error) {
	fmt.Println("CreatingFile...")

	conf = map[string]string{}

	file, err := os.Create(fileName + ".json")
	if err != nil {
		return conf, err
	}

	defer file.Close()
	var number int
	fmt.Println("Enter the number of parameter  :")
	fmt.Scan(&number)

	for i := 0; i < number; i++ {
		var name string
		var value string
		fmt.Println("---------------parameter ", i+1, " --------------- ")
		fmt.Print("name : ")
		fmt.Scan(&name)
		fmt.Print("value : ")
		fmt.Scan(&value)
		conf[name] = value

	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "	")
	err = encoder.Encode(conf)
	if err != nil {
		return conf, err
	}

	return conf, err
}

func ReadJSON(filePath string) (conf interface{}, err error) {
	fmt.Println("ReadingFiles...")

	file, err := os.Open(filePath + ".json")

	if err != nil {
		return CreateJSON(filePath)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&conf)
	if err != nil {
		return conf, err
	}

	return conf, err

}
