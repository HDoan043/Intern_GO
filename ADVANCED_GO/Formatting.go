package main

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

type Mark struct {
	Math   float32 `json:"math" yaml:"math"`
	Physic float32 `json:"physic" yaml:"physic"`
}

type Student struct {
	Name string `json:"name" yaml:"name"`

	Id int `json:"id" yaml:"id"`

	Mark Mark `json:"mark" yaml:"mark"`
}

func main() {
	// JSON unmarshal and marshal
	rawJSON := `{"name": "A", "id": 1, "mark": {"math": 10.0, "physic": 9.8}}`
	JSONMessage := []byte(rawJSON)

	var student Student
	json.Unmarshal(JSONMessage, &student)
	fmt.Printf("%+v\n", student)

	reverseJSON, _ := json.Marshal(student)
	fmt.Printf("%+v\n", string(reverseJSON))

	// YAML
	rawYAML := `
	"name": "A" 
	"id": 1
	"mark": 
		"math": 10.0
		"physic": 9.8`
	YAMLMessage := []byte(rawYAML)
	yaml.Unmarshal(YAMLMessage, &student)
	fmt.Printf("%+v\n", student)
	reverseYAML, _ := yaml.Marshal(student)
	fmt.Printf("%+v", string(reverseYAML))
}
