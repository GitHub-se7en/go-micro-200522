package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var DB *mongo.Database

func init() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		log.Fatalln("client err",err)
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalln("connect err",err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalln("ping err is ",err)
	}
	DB = client.Database("micro")

	//ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	//collection := DB.Collection("numbers")
	//res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	//if err != nil {
	//	log.Println(err)
	//}
	//id := res.InsertedID
	//log.Println(id)

}
