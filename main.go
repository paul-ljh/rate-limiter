package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Route struct {
	Endpoint  string `json:"endpoint"`
	Burst     int    `json:"burst"`
	Sustained int    `json:"sustained"`
}

func main() {
	configFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully opened config.json")
	defer configFile.Close()

	configBytes, err := io.ReadAll(configFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully read config.json")

	config := []Route{}
	if err := json.Unmarshal(configBytes, &config); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully loaded config.json")
}
