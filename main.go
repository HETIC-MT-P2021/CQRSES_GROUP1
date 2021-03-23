package main

import (
	"fmt"
	"os"

	"github.com/couchbase/gocb"
	"github.com/joho/godotenv"

	"log"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP1/controllers"
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

	// time.Sleep(15 * time.Second)

	cluster, _ := gocb.Connect("couchbase://localhost")
	clusterErr := cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: os.Getenv("COUCH_USER"),
		Password: os.Getenv("COUCH_PASSWORD"),
	})

	if clusterErr != nil {
		log.Fatalf("Error authenticating with couchbase %v", err)
	} else {
		fmt.Println("Connected to Couchbase")

		// clusterManager := cluster.Manager(os.Getenv("COUCH_USER"), os.Getenv("COUCH_PASSWORD"))

		// err1 := clusterManager.InsertBucket(&gocb.BucketSettings{
		// 	FlushEnabled:  false,
		// 	IndexReplicas: false,
		// 	Name:          "event-store",
		// 	Password:      "password",
		// 	Quota:         100,
		// 	Replicas:      1,
		// 	Type:          0,
		// })
		// if err1 != nil {
		// 	log.Fatalf("Error when inserting the bucket %v", err1)
		// }

		// err2 := clusterManager.InsertBucket(&gocb.BucketSettings{
		// 	FlushEnabled:  false,
		// 	IndexReplicas: false,
		// 	Name:          "read-models",
		// 	Password:      "password",
		// 	Quota:         100,
		// 	Replicas:      1,
		// 	Type:          0,
		// })
		// if err2 != nil {
		// 	log.Fatalf("Error when inserting the bucket %v", err1)
		// }
	}

	// Don't seed data here. Use docker-entrypoint instead
	// seed.Load(server.DB)

	server.Run(":8080")
}
