package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello")

	err := godotenv.Load()
	if err != nil {
		panic(".env file not found")
	}
}
