package api

import (
	"context"
	"fmt"
	"encoding/json"

	"net/http"

	"github.com/gorilla/websocket"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var upgrader = websocket.Upgrader {
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

type Message struct{
	Msg string `json:"msg"`
}

func ChatServerLogger(loggingCollection *mongo.Collection, ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}

		defer conn.Close()

		for {
			// Receive message
			messageType, message, err := conn.ReadMessage()	
			// send message
			if err != nil {
				// handle error
				panic(err)
			}

			err = conn.WriteMessage(messageType, message)
			if err != nil {
				panic(err)
			}

			_, err = loggingCollection.InsertOne(ctx, bson.M{"msg": string(message)})
			if err != nil {
				panic(err)
			}
		}
	}
}

func GetLogs(loggingCollection *mongo.Collection, ctx context.Context) func (w http.ResponseWriter, r *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		var logs []bson.M
		logs = queryLogs(loggingCollection, ctx)
		var messages []Message
		messages = parseLogs(logs)
		messagesJson, err := json.Marshal(messages)
		if err != nil {
			fmt.Println("ERROOR 4",err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(messagesJson)
	}
}

func queryLogs(loggingCollection *mongo.Collection, ctx context.Context) []bson.M{
	var logs []bson.M
	cursor, err := loggingCollection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("ERROOR 1",err)
	}
	if err = cursor.All(ctx, &logs); err != nil {
		fmt.Println("ERROOR 2",err)
	}
	if err != nil {
		fmt.Println("ERROOR 3",err)
	}
	return logs
}

func parseLogs(logs []bson.M) []Message{
	messages := []Message{}
	for _, v := range logs {
		str, ok := v["msg"].(string)
		if !ok {
			fmt.Printf("ERROR: not a string -> %#v\n", v["msg"])
			continue
		}
		var message = Message{
			Msg: str,
		}
		messages = append(messages, message)
	}
	return messages
}

func insertLogRecord(loggingCollection *mongo.Collection, ctx context.Context, message string){
	loggingCollection.InsertOne(ctx, bson.M{"msg": string(message)})
}

func deleteLogRecords(loggingCollection *mongo.Collection, ctx context.Context){
	loggingCollection.DeleteMany(ctx, bson.M{})
}