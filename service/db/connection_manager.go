package db

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

config := &firebase.Config{
        StorageBucket: "<BUCKET_NAME>.appspot.com",
}

// ConnectionManager for managing all db related connections
type ConnectionManager struct{}

// Init initialize new db connection manager
func (c *ConnectionManager) Init() {
	config := &firebase.Config{
		StorageBucket: "<BUCKET_NAME>.appspot.com",
	}
	opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}
}
