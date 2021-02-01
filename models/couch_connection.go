package models

import (
	"fmt"
	"time"

	"github.com/couchbase/gocb"
)

var eventBucket *gocb.Bucket
var readBucket *gocb.Bucket

func ConnectToCouchBase(host, user, password string) {
	fmt.Println("Waiting for couchBase")
	time.Sleep(12 * time.Second)
	fmt.Println("Connecting to CouchBase")
	cluster, couchConErr := gocb.Connect(fmt.Sprintf("couchbase://%s", host))

	numberOfTest := 0

	for couchConErr != nil && numberOfTest < 5 {
		fmt.Println(couchConErr)
		fmt.Println("Connection to CouchBase did not succeed, new try")

		time.Sleep(5 * time.Second)
		cluster, couchConErr = gocb.Connect(fmt.Sprintf("couchbase://%s", host))

		numberOfTest++
	}

	authErr := cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: user,
		Password: password,
	})

	if authErr != nil {
		fmt.Println("CouchBase Auth error")
		panic(authErr)
	}

	tempEventBucket, eventConErr := cluster.OpenBucket("event-store", "")

	if eventConErr != nil {
		fmt.Println("Couch Event bucket error")

		panic(eventConErr)
	}

	tempReadBucket, readConErr := cluster.OpenBucket("read-models", "")

	if readConErr != nil {
		fmt.Println("Couch Read bucket error")

		panic(readConErr)
	}

	if readConErr != nil {
		fmt.Println("Couch Event bucket error")
		panic(readConErr)
	}

	fmt.Println("CouchBase successfully connected!")

	eventBucket = tempEventBucket
	readBucket = tempReadBucket
}
