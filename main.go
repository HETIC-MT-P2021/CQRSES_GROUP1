package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"log"
	"os"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP1/controllers"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP1/seed"
)

var server = controllers.Server{}

func main() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	seed.Load(server.DB)

	server.Run(":8080")

}