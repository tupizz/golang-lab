package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func ConverteEmJSON(obj any) string {
	JSON, error := json.MarshalIndent(obj, " ", "  ")
	if error != nil {
		log.Fatal("erro em conversao")
	}

	return bytes.NewBuffer(JSON).String()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dataBytes, err := ioutil.ReadFile("/Users/tupizz/Desktop/tupizz/golang/golang-lab/24-json-unmarshal/exemplo.json")
	check(err)

	fmt.Println(string(dataBytes))

	var dev developer
	if erro := json.Unmarshal(dataBytes, &dev); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(dev.toJSON())

	cachorroEmJson := `{ "name": "Hulk", "raca": "Chihauha" }`
	cachorroMap := make(map[string]string)
	if erro := json.Unmarshal([]byte(cachorroEmJson), &cachorroMap); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroMap)

}
