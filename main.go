package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	connect := os.Getenv("MONGOSTRING")
	if connect == "" {
		log.Fatal(errors.New("no connect str"))
	}

	clientOptions := options.Client().ApplyURI(connect)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	userUpd(client.Database("i9").Collection("user"))
	strWoUpd(client.Database("i9").Collection("stretchworkout"))
	WoUpd(client.Database("i9").Collection("workout"))

	fmt.Printf("\n\nDONE\n")

}

func userUpd(collection *mongo.Collection) {
	var results []User
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	for _, item := range results {
		filter := bson.M{"_id": item.ID}
		update := bson.M{
			"$set": bson.M{
				"completed": 0,
				"badges":    []string{},
			},
		}
		_, err := collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("All users updated successfully!")
}

func strWoUpd(collection *mongo.Collection) {
	var results []StretchWorkout
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	for _, item := range results {
		filter := bson.M{"_id": item.ID}
		update := bson.M{
			"$set": bson.M{
				"laststarted": item.Created,
				"datelist":    []primitive.DateTime{item.Created},
				"pinned":      false,
				"startedct":   0,
			},
		}
		_, err := collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("All strwos updated successfully!")
}

func WoUpd(collection *mongo.Collection) {
	var results []Workout
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	for _, item := range results {

		item.LastStarted = item.Created
		item.DateList = []primitive.DateTime{item.Created}
		item.IsPinned = false
		item.AvgRating = float32(-1)
		item.AvgFaves = float32(-1)
		item.LastRating = -1
		item.LastFaves = -1
		item.RatedCount = 0
		item.StartedCount = 0

		for i := range item.Exercises {
			item.Exercises[i].AvgRating = float32(-1)
			item.Exercises[i].AvgFaves = float32(-1)
		}

		filter := bson.M{"_id": item.ID}
		update := bson.M{
			"$set": item,
		}

		_, err := collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("All wos updated successfully!")
}
