package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
	"trails/adder/model"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	JSON_FILE = "workout.json"
	MONGO_URI = "mongodb://localhost:27017"
	DB_NAME   = "trails"
	COLL_NAME = "workouts"
)

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {

	// open file
	file, err := os.Open(JSON_FILE)
	handleError(err)
	defer file.Close()

	// read file
	raw, err := ioutil.ReadAll(file)
	handleError(err)

	// unmarshal file
	var data model.Workout
	err = json.Unmarshal(raw, &data)
	handleError(err)

	// open db
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MONGO_URI))
	handleError(err)
	defer client.Disconnect(ctx)
	coll := client.Database(DB_NAME).Collection(COLL_NAME)

	// insert in db
	_, err = coll.InsertOne(ctx, data)
	handleError(err)

	fmt.Println("ok")
}
