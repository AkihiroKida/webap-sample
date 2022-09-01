package main

import (
	"fmt"
	"log"
	"os"

	"go-app/driver"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env not found")
	}

	log.Println("Server running...")
	driver.Serve(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
