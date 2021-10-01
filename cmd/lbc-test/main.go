package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/monkeydioude/lbc-test/internal/handler"
	"github.com/monkeydioude/lbc-test/pkg/bbolt"
	"github.com/monkeydioude/lbc-test/pkg/router"
)

const (
	defaultPort = 8004
)

// RuntimeServerConfig struct defines the config http server will use at launch
type RuntimeServerConfig struct {
	Addr   string
	DBPath string
}

// buildConfig returns a RuntimeServerConfig struct built around
// env vars and default values
func buildConfig() RuntimeServerConfig {
	rtc := RuntimeServerConfig{
		Addr:   fmt.Sprintf(":%d", defaultPort),
		DBPath: "./db.db",
	}

	if port := os.Getenv("PORT"); port == "" {
		log.Printf("[WARN] \"PORT\" env var not provided. Using default value \"%d\"\n", defaultPort)
	} else {
		rtc.Addr = fmt.Sprintf(":%s", port)
	}

	if dbPath := os.Getenv("DB_PATH"); dbPath == "" {
		log.Printf("[WARN] \"DB_PATH\" env var not provided. Using default value \"%s\"", rtc.DBPath)
	} else {
		rtc.DBPath = dbPath
	}

	return rtc
}

// main likes to hold and other hand
func main() {
	rtConf := buildConfig()

	bbolt.Open(rtConf.DBPath, 0666, nil)
	defer bbolt.Close()

	router := router.New()
	router.Post("/fizz-buzz", handler.FizzBuzzTestHandler)
	router.Get("/fizz-buzz/stats", handler.FizzBuzzStatsHandler)
	log.Printf("[INFO] starting server %+v\n", rtConf)
	http.ListenAndServe(rtConf.Addr, router)
}
