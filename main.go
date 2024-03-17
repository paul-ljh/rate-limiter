package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Route struct {
	Endpoint  string `json:"endpoint"`
	Burst     int    `json:"burst"`
	Sustained int    `json:"sustained"`
}

func loadConfig(config *[]Route) {
	configFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Successfully opened config.json")
	defer configFile.Close()

	configBytes, err := io.ReadAll(configFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Successfully read config.json")

	if err := json.Unmarshal(configBytes, &config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Successfully loaded config.json")
}

func take(w http.ResponseWriter, r *http.Request) {
	endpoint := r.PostFormValue("endpoint")
	fmt.Printf("Checking rate limit for endpoint: %s\n", endpoint)

	io.WriteString(w, "Check my boi\n")
}

func main() {
	config := []Route{}
	loadConfig(&config)

	mux := http.NewServeMux()
	mux.HandleFunc("/take", take)

	server := &http.Server{
		Addr:    ":3333",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server closed")
	} else if err != nil {
		fmt.Printf("error listening for server: %s\n", err)
	}
}
