package libs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"slices"
	"strings"
)

func LoadOperations(json_file string) (string, []Operation) {
	var results []Operation
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

	url := result["url"].(string)
	var root = result["tests"].([]any)[0].(map[string]interface{})
	var commands = root["commands"].([]any)

	for i, _ := range commands {
		var command = commands[i].(map[string]interface{})

		var oper Operation
		oper.Action = command["command"].(string)

		doubleMethods := []string{"click", "type", "sendKeys"}
		if slices.Contains(doubleMethods, oper.Action) {
			content := strings.Split(command["target"].(string), "=")
			oper.Target = content[0]
			oper.Tag = content[1]
			oper.Value = command["value"].(string)

			// fmt.Println(oper)
			results = append(results, oper)
		}

		singleMethods := []string{"open"}
		if slices.Contains(singleMethods, oper.Action) {
			oper.Target = command["target"].(string)
			oper.Value = command["value"].(string)

			// fmt.Println(oper)
			results = append(results, oper)
		}
	}

	return url, results
}

func LoadJsonFromFile(filename string) map[string]int {
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]int
	json.Unmarshal([]byte(byteValue), &result)
	return result
}
