package database


import (
	"fmt"
	"context"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
)

func DBinstance() *mongo.Client {
	MongoDB := "mongodb://localhost:27017" 
	fmt.Println(MongoDB) 

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDB))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

var Client *mongo.Client = DBinstance()


func OpenCollection (client *mongo.Client, collectionName string) *mongo.Collection {

var collection *mongo.Collection = client.Database("restuarant").Collection(collectionName)
fmt.Println("Collection is opened")

return collection
} 

