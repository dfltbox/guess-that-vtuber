package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func pushtodb() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("CONNECTIONURL")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	file, err := os.OpenFile("formatted.json", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var vtubers []VtuberStruct
	if err := json.Unmarshal(byteValue, &vtubers); err != nil {
		log.Fatal(err)
	}

	collection := client.Database("vtuberdb").Collection("vtubers")

	var documents []interface{}
	for _, vtuber := range vtubers {
		documents = append(documents, vtuber)
	}

	result, err := collection.InsertMany(ctx, documents)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Added %v vtubers\n", len(result.InsertedIDs))
}
