package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/monkeydioude/lbc-test/internal/handler"
	"github.com/monkeydioude/lbc-test/pkg/router"
)

const (
	defaultPort = 8004
)

// RuntimeServerConfig struct defines the config http server will use at launch
type RuntimeServerConfig struct {
	Addr string
}

// buildConfig returns a RuntimeServerConfig struct built around
// env vars and default values
func buildConfig() RuntimeServerConfig {
	rtc := RuntimeServerConfig{}

	if port := os.Getenv("PORT"); port == "" {
		log.Printf("[WARN] \"PORT\" env var not provided. Using default value \"%d\"\n", defaultPort)
		rtc.Addr = fmt.Sprintf(":%d", defaultPort)
	} else {
		rtc.Addr = fmt.Sprintf(":%s", port)
	}

	return rtc
}

// main likes to hold and other hand
func main() {
	rtConf := buildConfig()
	router := router.New()
	router.Post("/fizz-buzz/test", handler.FizzBuzzTestHandler)
	http.ListenAndServe(rtConf.Addr, router)
}
