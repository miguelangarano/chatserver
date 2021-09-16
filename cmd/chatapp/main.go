package main

import (
	"context"
	"time"
	"log"

	"net/http"
	"jrtlabs/chatapp/src/api"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 60*5* time.Second)
	defer cancel()
	
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:password@mongo:27017"))
	defer client.Disconnect(ctx)
	if err != nil {
		panic(err)
	}

	db := client.Database("chatappLogging")
	loggingCollection := db.Collection("logs")
	http.HandleFunc("/socket", api.ChatServerLogger(loggingCollection, ctx))
	http.HandleFunc("/logs", api.GetLogs(loggingCollection, ctx))

	log.Fatal(http.ListenAndServe("0.0.0.0:3000", nil))
}
