package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://khushal:Khushal4989@cluster0.xirtede.mongodb.net/?retryWrites=true&w=majority"
const dbName = "garage"
const colName = "carList"

var collection *mongo.Collection

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	// collection instance

	fmt.Println("Collection instance is ready")
}
