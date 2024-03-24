package main

import (
	"fmt"
	"time"
)

func initializeTokenRefiller(sustain int, rate *Rate) {
	d := calculateTickerDuration(sustain)
	ticker := time.NewTicker(time.Duration(d.Nanoseconds()))
	defer ticker.Stop()

	for {
		<-ticker.C
		rate.Replenish()
	}
}

func calculateTickerDuration(sustain int) time.Duration {
	t, _ := time.ParseDuration(fmt.Sprintf("%fs", float64(60)/float64(sustain)))
	return t
}
