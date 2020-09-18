package main

import (
	"github.com/joho/godotenv"
	"go_training/initializer"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	initializer.Init()
}