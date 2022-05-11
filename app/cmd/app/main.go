package main

import (
	"log"
	"prod/app/internal/config"
)

func main() {
	log.Print("config initializing")
	cfg := config.GetConfig()

	log.Print("config initializing")
}