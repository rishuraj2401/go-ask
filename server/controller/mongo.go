package mong

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rishuraj2401/quest/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"log"
	"net/http"
	"strconv"
)

func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

func InsertOne(quest model.Question) {
	data, err := collection.InsertOne(context.Background(), quest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted successfully", data)
}
func Inserted(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is router 5")
	setupCorsResponse(&w, r)
	w.Header().Set("content-type", "application/json")

	data, _ := io.ReadAll(r.Body)

	var question model.Question
	if question.ID.IsZero() {
		question.ID = primitive.NewObjectID()
	}
	err := json.Unmarshal(data, &question)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("this is body", question, string(data))
	InsertOne(question)
	w.Write([]byte("Inseted succesfully"))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is router 0")
	w.Header().Set("content-type", "application/json")
	w.Write([]byte("Welcome succesfully"))
}

func update(qId string, ans model.Answer, user string) {
	id, _ := primitive.ObjectIDFromHex(qId)
	ans.AnsBy = user
	filter := bson.M{"_id": id}
	fmt.Println("something is here", id)
	update := bson.M{"$push": bson.M{
		"answer": ans}}
	uData, _ := collection.UpdateOne(context.Background(), filter, update)
	fmt.Println("answer is here", uData)

}


func AddAnswer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is router 3")
	w.Header().Set("content-type", "application/json")
	param := mux.Vars(r)
	user := param["user"]
	data, _ := io.ReadAll(r.Body)
	var ans model.Answer
	err := json.Unmarshal(data, &ans)
	if ans.ID.IsZero() {
		ans.ID = primitive.NewObjectID()
	}
	if err != nil {
		log.Fatal(err)
	}

	update(param["id"], ans, user)
	w.Write([]byte("hey over"))
	fmt.Println("answer is updated", param["id"])
}


func GetQuestionsByPage(pageNumber int) ([]model.Question, error) {
	skip := (pageNumber - 1) * 5

	// Define options for pagination
	options := options.Find().SetLimit(int64(5)).SetSkip(int64(skip))

	// Define a filter to match all documents
	filter := bson.D{{}}

	// Perform the query
	cursor, err := collection.Find(context.TODO(), filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var questions []model.Question
	for cursor.Next(context.TODO()) {
		var question model.Question
		err := cursor.Decode(&question)
		if err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}
	fmt.Println(questions)
	return questions, nil
}


func GetQuestion(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	page := param["page"]
	pageNum, _ := strconv.Atoi(page)
	questions, err := GetQuestionsByPage(pageNum)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Convert questions to JSON
	jsonData, err := json.Marshal(questions)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(jsonData)

}
