package main

import (
	"context"

	config "github.com/joycezemitchell/product-grpc-api/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	config.Collection.Find(context.Background(), primitive.D{{}})
}
