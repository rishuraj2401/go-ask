package mong

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

const dbName string = "dbQ"
const col string = "Questions"
const users string = "Users"

var collection *mongo.Collection
var userCol *mongo.Collection

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error in loading env")
	}
	url := os.Getenv("URL")
	// const url string = "mongodb+srv://rishuraj2401:Rishu%402002@cluster0.twrql.mongodb.net"
	clientOpt := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.TODO(), clientOpt)
	if err != nil {
		log.Fatal(err)
	}
	index := mongo.IndexModel{
		Keys: bson.M{
			"questions": "text",
		},
	}
	index1 := mongo.IndexModel{
		Keys: bson.M{
			"name": "text",
		},
	}
	// _, err = client.Database(dbName).Collection(col).Indexes().CreateOne(context.TODO(), index)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println("Db is connected")
	userCol = client.Database(dbName).Collection(users)

	collection = client.Database(dbName).Collection(col)
	userCol.Indexes().CreateOne(context.TODO(), index1)

	collection.Indexes().CreateOne(context.TODO(), index)
	fmt.Println("Coleection instance is ready", userCol, collection)
	// Add this code after creating the collection

}
