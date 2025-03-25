package libs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func LoadJsonFromFile(json_file string) string {
	// plan, _ := ioutil.ReadFile(json_file)
	// var data interface{}
	// err := json.Unmarshal(plan, &data)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return "{}"
	// }
	// return data
	// Open our jsonFile
	jsonFile, err := os.Open(json_file)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	var root = result["tests"].([]any)[0].(map[string]interface{})
	var commands = root["commands"].([]any)
	var command = commands[0].(map[string]interface{})
	fmt.Println(command["command"])
	fmt.Println(command["target"])
	fmt.Println(command["id"])

	return "{}"
}
