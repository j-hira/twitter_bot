package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	fmt.Println(apiKey)
}
