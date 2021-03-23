package models

import (
	"fmt"

	"github.com/couchbase/gocb"
)

var ReadBucket *gocb.Bucket

func ConnectToCouchBase(host, user, password string) {
	fmt.Println(host)
	fmt.Println("Connecting to CouchBase")

	cluster, couchConErr := gocb.Connect(fmt.Sprintf("couchbase://%s", host))

	if couchConErr != nil {
		fmt.Println("CouchBase Connect error")
		panic(couchConErr)
	}

	authErr := cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: user,
		Password: password,
	})

	if authErr != nil {
		fmt.Println("CouchBase Auth error")
		panic(authErr)
	}

	tempReadBucket, readConErr := cluster.OpenBucket("read-models", "")

	if readConErr != nil {
		fmt.Println("Couch Read bucket error")

		panic(readConErr)
	}

	fmt.Println("CouchBase successfully connected!")

	ReadBucket = tempReadBucket
}
