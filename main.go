package main

import (
	"fmt"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Starting Bot")

	env, _ := godotenv.Read(".env")

	dbPort, dbErr := strconv.ParseInt(env["DB_PORT"], 10, 64)

	if dbErr != nil {
		panic(dbErr)
	}

	models.ConnectToDB(env["DB_HOST"], env["DB_NAME"], env["DB_USER"], env["DB_PASSWORD"], dbPort)


}
