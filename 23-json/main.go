package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type developer struct {
	Name   string   `json:"name"`
	Hobby  string   `json:"hobby"`
	Age    int8     `json:"age"`
	Skills []string `json:"skills"`
}

func (d developer) toJSON() string {
	return ConverteEmJSON(d)
}

func ConverteEmJSON(obj interface{}) string {
	JSON, error := json.MarshalIndent(obj, " ", "  ")
	if error != nil {
		log.Fatal("erro em conversao")
	}

	return bytes.NewBuffer(JSON).String()
}

func main() {
	d := developer{
		Name:  "tadeu",
		Hobby: "cross fitz",
		Age:   27,
		Skills: []string{
			"golang",
			"nodejs",
			"devops",
			"kubernetes",
			"java",
			"react",
			"vue",
		},
	}

	fmt.Println(d.toJSON())

}
