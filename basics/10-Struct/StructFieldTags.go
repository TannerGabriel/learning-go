package main

import (
	"encoding/json"
	"fmt"
)

type Dog struct {
	Name   string `json:"name"`
	Weight string `json:"weight"`
}

func main() {
	json_string := `
    {
        "name": "Rocky",
        "weight": "45"
    }`

	rocky := new(Dog)
	json.Unmarshal([]byte(json_string), rocky)
	fmt.Println(rocky)

	spot := new(Dog)
	spot.Name = "Spot"
	spot.Weight = "20"
	jsonStr, _ := json.Marshal(spot)
	fmt.Printf("%s\n", jsonStr)
}
