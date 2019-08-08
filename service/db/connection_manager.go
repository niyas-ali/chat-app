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
type ConnectionManager struct{
   client struct
}

// Init initialize new db connection manager
func (c *ConnectionManager) Init() {
	
  // Use a service account
  ctx := context.Background()
  sa := option.WithCredentialsFile("configs/configuration.json")
  app, err := firebase.NewApp(ctx, nil, sa)
  if err != nil {
    log.Fatalln(err)
  }

  client, err := app.Firestore(ctx)
  if err != nil {
    log.Fatalln(err)
  }
  defer client.Close()
}

func (c *ConnectionManager) Add(){
  _, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
       
})
if err != nil {
        log.Fatalf("Failed adding record: %v", err)
}
}

func (c *ConnectionManager) Read(){
  iter := client.Collection("users").Documents(ctx)
for {
        doc, err := iter.Next()
        if err == iterator.Done {
                break
        }
        if err != nil {
                log.Fatalf("Failed to iterate: %v", err)
        }
        fmt.Println(doc.Data())
}
}
