package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Percentages struct {
	One   float64 `json:"one"`
	Two   float64 `json:"two"`
	Three float64 `json:"three"`
}

type Params struct {
	Pages    string        `json:"page"`
	Words    []string      `json:"words"`
	Percents Percentages   `json:"percentages"`
	Special  []string      `json:"special"`
	Extra    []interface{} `json:"extraSpecial"`
}

func main() {

	var params Params
	b, _ := os.ReadFile("./example.json")
	json.Unmarshal(b, &params)

	fmt.Println(params)
}
