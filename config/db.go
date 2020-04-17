package config

import "fmt"

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// var collection *mongo.Collection

// func init() {

// 	// connect to MongoDB
// 	fmt.Println("Connecting to MongoDB")
// 	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = client.Connect(context.TODO())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	collection = client.Database("allyshop").Collection("products")
// }

// TestPrint implements
func TestPrint() {
	fmt.Println("TestPrint")
}
