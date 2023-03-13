package connectdb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Function used to connect with database
func Connect() *mongo.Collection {

	// connecting with the database
	sess, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected with MongoDB!")

	userCollection := sess.Database("moviesdb").Collection("movies")

	if err != nil {
		log.Fatal(err)
	}

	return userCollection
}
