package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/yash-278/gosync/internal/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	routes.StartServer()
}
