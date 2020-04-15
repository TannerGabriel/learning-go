package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Name  string
	Age   int
	Email string
}

func main() {
	// Set client options - Port, URL
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Get a handle for your collection
	collection := client.Database("test").Collection("users")

	// Some dummy data to add to the Database
	john := User{"John", 20, "johndoe@test.com"}
	maria := User{"Maria", 10, "maria@test.com"}
	brock := User{"Brock", 15, "brock@example.com"}

	// Insert a single document
	insertUser(collection, john)

	// Insert multiple documents
	users := []interface{}{maria, brock}
	insertMultipleUsers(collection, users)

	// Update a document
	updateUser(collection, maria, "John")

	// Find a single document
	findUser(collection, "Maria")

	// Finding multiple documents
	findAllUsers(collection)

	// Delete all the documents in the collection
	deleteAllUsers(collection)

	// Close the connection once no longer needed
	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection to MongoDB closed.")
	}

}

// Insert a document into the database
func insertUser(collection *mongo.Collection, user User) {
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

// Insert multiple documents into the database
func insertMultipleUsers(collection *mongo.Collection, users []interface{}) {
	insertManyResult, err := collection.InsertMany(context.TODO(), users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
}

// Update a document in the database
func updateUser(collection *mongo.Collection, user User, userName string) {
	filter := bson.D{{"name", userName}}

	update := bson.D{
		{"$set", bson.D{{"user", user}}},
		{"$currentDate", bson.D{{"modifiedAt", true}}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}

// Find a document of the database
func findUser(collection *mongo.Collection, userName string) User {
	filter := bson.D{{"name", userName}}

	var result User

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)

	return result
}

// Find all documents of a collection
func findAllUsers(collection *mongo.Collection) {
	findOptions := options.Find()
	findOptions.SetLimit(2)

	var results []*User

	// Finding multiple documents returns a cursor
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through the cursor
	for cur.Next(context.TODO()) {
		var elem User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
}

// Delete a document of the database
func deleteAllUsers(collection *mongo.Collection) {
	deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v documents in the users collection\n", deleteResult.DeletedCount)
}
