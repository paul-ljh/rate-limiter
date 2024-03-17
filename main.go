package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

type ctxKey string

const configCtxKey ctxKey = "config"

type RateConfig struct {
	Burst     int `json:"burst"`
	Sustained int `json:"sustained"`
}

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

func transformConfig(config *[]Route, configMap *map[string]RateConfig) {
	for _, route := range *config {
		(*configMap)[route.Endpoint] = RateConfig{
			Burst:     route.Burst,
			Sustained: route.Sustained,
		}
	}
}

func take(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	config := *(ctx.Value(configCtxKey).(*map[string]RateConfig))

	endpoint := r.PostFormValue("endpoint")
	if endpoint == "" {
		io.WriteString(w, fmt.Sprintln("You must provide endpoint in your request body"))
	} else if _, ok := config[endpoint]; !ok {
		io.WriteString(w, fmt.Sprintln("Invalid endpoint in your request body"))
	} else {
		io.WriteString(w,
			fmt.Sprintf(
				"The rate limit config for endpoint %s is: {Burst: %d, Sustained: %d}\n",
				endpoint,
				config[endpoint].Burst,
				config[endpoint].Sustained,
			),
		)
	}
	fmt.Printf("Checking rate limit for endpoint: %s\n", endpoint)
}

func main() {
	config := []Route{}
	loadConfig(&config)

	configMap := map[string]RateConfig{}
	transformConfig(&config, &configMap)

	mux := http.NewServeMux()
	mux.HandleFunc("/take", take)

	ctx := context.Background()
	server := &http.Server{
		Addr:    ":3333",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, configCtxKey, &configMap)
			return ctx
		},
	}

	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server closed")
	} else if err != nil {
		fmt.Printf("error listening for server: %s\n", err)
	}
}
