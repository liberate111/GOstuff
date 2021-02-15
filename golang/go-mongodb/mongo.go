package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// Set client option
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, e := mongo.Connect(context.TODO(), clientOptions)
	CheckError(e)

	// Check the connection
	e = client.Ping(context.TODO(), nil)
	CheckError(e)

	fmt.Println("Connected to MongoDB!")

	// get collection as ref
	collection := client.Database("testdb").Collection("people")

	// Insert
	// john := Person{"John", 20}
	// jane := Person{"Jane", 21}
	// ben := Person{"Ben", 22}

	// _, e = collection.InsertOne(context.TODO(), john)
	// CheckError(e)

	// persons := []interface{}{jane, ben}
	// _, e = collection.InsertMany(context.TODO(), persons)
	// CheckError(e)

	// Update
	filter := bson.D{{"name", "John"}}
	// update := bson.D{
	// 	{"$set", bson.D{
	// 		{"age", 30},
	// 	}},
	// }

	// _, e = collection.UpdateOne(context.TODO(), filter, update)
	// CheckError(e)

	// Get data
	var res Person
	e = collection.FindOne(context.TODO(), filter).Decode(&res)
	fmt.Println(res)

	// Delete
	// _, e = collection.DeleteOne(context.TODO(), filter)
	// CheckError(e)

	// _, e = collection.DeleteMany(context.TODO(), bson.D{{}})
	// CheckError(e)

}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
