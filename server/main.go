package main

import (
	"fmt"

	"github.com/fr3fou/aumo/server/aumo"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello")
	a := aumo.New(aumo.Config{})

	fmt.Printf("%+v", a)
	err := godotenv.Load()
	if err != nil {
		panic(".env file not found")
	}
}
