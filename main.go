package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"log"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP1/controllers"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP1/models"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP1/producer"
)

var server = controllers.Server{}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"), os.Getenv("COUCH_HOST"), os.Getenv("COUCH_USER"), os.Getenv("COUCH_PASSWORD"))

	producer.ConnectToRabbit(
		os.Getenv("RABBIT_HOST"),
		os.Getenv("RABBIT_PORT"),
		os.Getenv("RABBIT_USER"),
		os.Getenv("RABBIT_PASSWORD"),
	)

	models.ConnectToCouchBase(os.Getenv("COUCH_HOST"), os.Getenv("COUCH_USER"), os.Getenv("COUCH_PASSWORD"))

	server.Run(":8080")
}
