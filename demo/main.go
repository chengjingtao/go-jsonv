package main

import (
	"encoding/json"
	"fmt"

	jsonv "github.com/chengjingtao/go-jsonv"
	"gopkg.in/yaml.v2"
)

type kv struct {
	Name  string
	Value jsonv.JsonV
}

func main() {
	demoYaml()
	demoJSON()
}

func demoJSON() {
	var data = `
	{
		"name": "the-name",
		"value": [
			"1",
			true,
			{
				"key1": "1",
				"key2": 2
			}
		]
	}`

	m := kv{}
	err := json.Unmarshal([]byte(data), &m)
	fmt.Printf("Unmarshal err: %#v\n", err)
	fmt.Printf("%#v\n", m)

	byts, err := json.MarshalIndent(m, "", "  ")
	fmt.Printf("err: %#v\n", err)
	fmt.Println(string(byts))
}

func demoYaml() {
	var data = `
name: "the-name"
value: 
  - "1"
  - true
  - key1: "1"
    key2: 2
`

	m := kv{}
	err := yaml.Unmarshal([]byte(data), &m)
	fmt.Printf("Unmarshal err: %#v\n", err)
	fmt.Printf("%#v\n", m)

	byts, err := yaml.Marshal(m)
	fmt.Printf("Marshal err: %#v\n", err)
	fmt.Println(string(byts))
}
