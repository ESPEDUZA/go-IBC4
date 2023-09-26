package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := dbConnect("mongodb://localhost:27017")
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}
	defer client.Disconnect(context.Background())

	resp, err := http.Get("https://api.kraken.com/0/public/Assets")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&result)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Save to MongoDB
	collection := client.Database("kraken").Collection("crypto")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	for _, value := range result {
		_, err := collection.InsertOne(ctx, value)
		if err != nil {
			fmt.Println("Error inserting into MongoDB:", err)
		}
	}

	fmt.Println("Data has been inserted into MongoDB")
}

func dbConnect(uri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	return client, nil
}
