package main

import (
	"context"
	"fmt"
	"log"

	client "github.com/codescalersinternships/DateTime-Client-Abdelrahman-Mahmoud/client"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	fmt.Println("Client created")

	myClient := client.NewClient()

	fmt.Println("HTTP server - GET /datetime  -> current date and time")
	returnedDateTime, err := myClient.GetHTTPDateTime(context.Background())

	if err != nil {
		log.Fatalf("error getting current date and time: %s", err)
	} else {
		fmt.Println(returnedDateTime)
	}

	fmt.Println("Gin server - GET /datetime  -> current date and time")
	returnedDateTime, err = myClient.GetGinDateTime(context.Background())

	if err != nil {
		log.Fatalf("error getting current date and time: %s", err)
	} else {
		fmt.Println(returnedDateTime)
	}

}
