package app

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewTodo(db *mongo.Database, data string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	podcastResult, err := db.Collection("todo").InsertOne(ctx, bson.D{
		{Key: "title", Value: data},
		{Key: "author", Value: data},
		{Key: "tags", Value: bson.A{"development", "programming", "coding"}},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(podcastResult)
}

func FetchTodo(db *mongo.Database) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := db.Collection("todo").Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var todo []bson.M
	if err = cursor.All(ctx, &todo); err != nil {
		log.Fatal(err)
	}
	fmt.Println(todo)
}
