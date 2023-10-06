package main

import (
	"log"
	"os"
)

func envChecks() {
	port, portExist := os.LookupEnv("PORT")

	if !portExist || port == "" {
		log.Fatal("PORT must be set in .env and not empty")
	}
}
