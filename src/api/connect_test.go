package api
import (
	"testing"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestParsing(t *testing.T){
	var logs []bson.M
	emptyResult := parseLogs(logs)
	if len(emptyResult)!=0 {
		t.Errorf("Test failed expected len is %v and received %v", "0", len(emptyResult))
	}else{
		t.Logf("Test succeded expected len is %v and received %v", "0", len(emptyResult))
	}
	logs = append(logs, bson.M{"_id":"613c3c690abd1b66338e3432", "msg":"ojossje"})
	logs = append(logs, bson.M{"_id":"613c3c690abd1b66338e3433", "msg":"what?"})
	result := parseLogs(logs)
	if len(result)!=2 {
		t.Errorf("Test failed expected len is %v and received %v", "2", len(result))
	}else{
		t.Logf("Test succeded expected len is %v and received %v", "2", len(result))
	}
}

func TestQueryLogs(t *testing.T){
	ctx, cancel := context.WithTimeout(context.Background(), 60*5* time.Second)
	defer cancel()
	
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:password@mongo:27017"))
	defer client.Disconnect(ctx)
	if err != nil {
		panic(err)
	}
	db := client.Database("chatappLogging")
	loggingCollection := db.Collection("logs")
	emptyResult := queryLogs(loggingCollection, ctx)
	if len(emptyResult)!=0 {
		t.Errorf("Test failed expected len is %v and received %v", "0", len(emptyResult))
	}else{
		t.Logf("Test succeded expected len is %v and received %v", "0", len(emptyResult))
	}
	insertLogRecord(loggingCollection, ctx, "My first message")
	insertLogRecord(loggingCollection, ctx, "My second message")
	result := queryLogs(loggingCollection, ctx)
	if len(result)!=2 {
		t.Errorf("Test failed expected len is %v and received %v", "2", len(result))
	}else{
		t.Logf("Test succeded expected len is %v and received %v", "2", len(result))
	}
	deleteLogRecords(loggingCollection, ctx)
}