package common

import "github.com/couchbase/gocb"

func Bucket() *gocb.Bucket {
	cluster, _ := gocb.Connect("couchbase://localhost")
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "admin",
		Password: "123456",
	})
	bucket, _ := cluster.OpenBucket("task", "")

	return bucket
}