package common

import (
	"github.com/couchbase/gocb"
	"os"
)

func Bucket() *gocb.Bucket {
	dbUsername, _ := os.LookupEnv("DB_USERNAME")
	dbPassword, _ := os.LookupEnv("DB_PASSWORD")

	cluster, _ := gocb.Connect("couchbase://localhost")
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: dbUsername,
		Password: dbPassword,
	})
	bucket, _ := cluster.OpenBucket("task", "")

	return bucket
}