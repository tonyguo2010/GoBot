package libs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"slices"
	"strings"
)

func LoadOperations(json_file string) []Operation {
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

	var root = result["tests"].([]any)[0].(map[string]interface{})
	var commands = root["commands"].([]any)

	for i, _ := range commands {
		var command = commands[i].(map[string]interface{})

		var oper Operation
		oper.Action = command["command"].(string)

		methods := []string{"click", "type"}
		if slices.Contains(methods, oper.Action) {
			content := strings.Split(command["target"].(string), "=")
			oper.Target = content[0]
			oper.Tag = content[1]
			oper.Value = command["value"].(string)

			// fmt.Println(oper)
			results = append(results, oper)
		}
	}

	return results
}
