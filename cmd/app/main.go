package main

import (
	"log"

	"github.com/joho/godotenv"
	"golang-restapi/internal/app/di"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("warning: unable to load .env file:", err)
	}

	container, err := di.NewContainer()
	if err != nil {
		log.Fatal(err)
	}

	if err := container.Run(); err != nil {
		log.Fatal(err)
	}
}
