package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ghodss/yaml"
)


type Standards struct {
	Family           string  `json:"family"`
	Name     string  `json:"name"`
	Description     string  `json:"desc"`
}


func main() {

	yamlFile, err := ioutil.ReadFile("/Users/gauravbang/Documents/meng/security-central/standards/nist-800-53-latest.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	output, err := yaml.YAMLToJSON(yamlFile)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	//fmt.Println(string(output))

	var result map[string]interface{}
	json.Unmarshal([]byte(output), &result)

	//fmt.Printf("Result: ", result["SI-7 (4)"].(map[string]interface{}))

	for key, value := range result {
		// Each value is an interface{} type, that is type asserted as a string
		fmt.Println(key)
		//fmt.Println(value.(map[string]interface{}))
		for k, v := range value.(map[string]interface{}) {
			fmt.Println(k, v)
		}
		//fmt.Println(value["family"].(map[string]interface{}))
		//fmt.Println(value["name"].(map[string]interface{}))
		break
	}

}
