package common

import (
	"github.com/couchbase/gocb"
	"os"
)

func Cluster() *gocb.Cluster {
	dbUsername, _ := os.LookupEnv("DB_USERNAME")
	dbPassword, _ := os.LookupEnv("DB_PASSWORD")

	cluster, _ := gocb.Connect("couchbase://localhost")
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: dbUsername,
		Password: dbPassword,
	})

	return cluster
}