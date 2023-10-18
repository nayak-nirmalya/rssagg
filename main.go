package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello World!")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the environment")
	}

	fmt.Println("Port:", port)
}
