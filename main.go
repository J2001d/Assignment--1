//  Implement a basic CRUD API: The intern should create a basic CRUD (Create, Read,
// 	Update, Delete) API using Go that allows a user to manage a list of items (e.g., books,
// 	movies, etc.). The API should include endpoints for creating, reading, updating, and deleting
// 	items.

// Name - Jhalak Dashora
// Mail - jhalakdashora01@gmail.com

package main

import (
	models "Assignment/Models"
	connectdb "Assignment/connectdb"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var usersCollection = connectdb.Connect()

// getting all movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// finding all data in the database using Find
	cursor, err := usersCollection.Find(context.TODO(), bson.D{})

	// checking if there is any error
	if err != nil {
		panic(err)
	}

	// convert the cursor result to bson
	var movies []bson.M

	// check for errors in the conversion
	if err = cursor.All(context.TODO(), &movies); err != nil {
		panic(err)
	}

	// displaying all the movies to console
	for _, movie := range movies {
		fmt.Println(movie)
	}

	// sendind data to postman/Frontend
	json.NewEncoder(w).Encode(movies)
}

// getting a particular movie with a id
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// taking that params from api request
	// where r is the request that user sends
	params := mux.Vars(r)

	var movie bson.M

	// string to primitive id
	id, _ := primitive.ObjectIDFromHex((params["id"]))

	// creating filter
	filter := bson.M{"_id": id}
	// check for errors in the finding
	// retrieving the first document that matches the filter
	if err := usersCollection.FindOne(context.TODO(), filter).Decode(&movie); err != nil {
		panic(err)
	}

	// sendind data to postman/Frontend
	json.NewEncoder(w).Encode(movie)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie models.Movie
	// decoding the json to struct
	_ = json.NewDecoder(r.Body).Decode(&movie)

	// insert the bson object using InsertOne()
	res, err := usersCollection.InsertOne(context.TODO(), movie)

	// check for errors in the insertion
	if err != nil {
		panic(err)
	}

	// sendind id to postman/Frontend
	json.NewEncoder(w).Encode(res)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// taking that params from api request
	// where r is the request that user sends
	params := mux.Vars(r)

	//Getting id from parameters
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var movie models.Movie

	// Creating filter
	filter := bson.M{"_id": id}

	// Read update model from body request
	_ = json.NewDecoder(r.Body).Decode(&movie)

	// preparing update model.
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "rating", Value: movie.Rating},
			{Key: "title", Value: movie.Title},
			{Key: "director", Value: movie.Director},
		}},
	}

	// finding and updating it in database
	err := usersCollection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&movie)

	if err != nil {
		panic(err)
	}

	movie.ID = id
	// sendind data to postman/Frontend
	json.NewEncoder(w).Encode(movie)
}

// deleting a movie with a id
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// taking that params from api request
	// where r is the request that user sends
	params := mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex((params["id"]))

	//creating filter
	filter := bson.M{"_id": id}

	// delete the first movie that match the id
	res, err := usersCollection.DeleteOne(context.TODO(), filter)

	// checking for errors in the deleting
	if err != nil {
		panic(err)
	}
	// sendind result to postman/Frontend
	json.NewEncoder(w).Encode(res)
}

func main() {
	// creating routes using gorilla mux
	route := mux.NewRouter()

	route.HandleFunc("/movies", getMovies).Methods("GET")

	route.HandleFunc("/movies", createMovie).Methods("POST")

	route.HandleFunc("/movies/{id}", getMovie).Methods("GET")

	route.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")

	route.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Staring server at PORT 4000")
	// listening to server
	log.Fatal(http.ListenAndServe(":4000", route))

}
