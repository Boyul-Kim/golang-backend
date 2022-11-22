package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongo_db struct {
	connection *mongo.Client
	ctx        context.Context
}

type User struct {
	Name string
}

func setupDB() mongo_db {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	databases, err := client.ListDatabaseNames(ctx, bson.M{})

	fmt.Println("DATABASES", databases)

	dbConnection := mongo_db{
		connection: client,
		ctx:        ctx,
	}

	collection := client.Database("cofi-lite").Collection("users")
	fmt.Println("Collection type:", reflect.TypeOf(collection))

	// Find all
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var results User
		if err = cursor.Decode(&results); err != nil {
			log.Fatal(err)
		}
		fmt.Println("RESULTS", results)
	}

	// var results User
	// err = collection.FindOne(ctx, bson.D{{"name", "saitama"}}).Decode(&results)
	// if err != nil {
	// 	fmt.Println("Error calling FindOne():", err)
	// 	os.Exit(1)
	// } else {
	// 	fmt.Println("FindOne() result:", results.Name)
	// }

	return dbConnection
}

func (m mongo_db) getUser() {
	collection := m.connection.Database("cofi-lite").Collection("users")
	fmt.Println("Collection type:", reflect.TypeOf(collection))

	var results User

	err := collection.FindOne(context.TODO(), bson.D{}).Decode(&results)

	if err != nil {
		fmt.Println("Error calling FindOne():", err)
		os.Exit(1)
	} else {
		fmt.Println("FindOne() result:", results)
		fmt.Println("FindOne() Name:", results.Name)
	}

	fmt.Println("RESULTS", results)
}
