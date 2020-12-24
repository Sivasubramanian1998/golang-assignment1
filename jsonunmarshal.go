package main

import (
	"encoding/json"
	"fmt"
)

type SensorReading struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
}

func main() {
	jsonString := `{"name": "battery sensor", "capacity": 40}`

	var reading SensorReading
	err := json.Unmarshal([]byte(jsonString), &reading)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", reading)
}
