package mong

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rishuraj2401/quest/model"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)
const pageSize=5

func SearchQuestionsByPage(pageNumber int, searchTerm string) ([]model.Question, error) {
	skip := (pageNumber - 1) * pageSize

	options := options.Find().SetLimit(int64(pageSize)).SetSkip(int64(skip))

	// Creating a filter to search for questions containing the searchTerm
	// filter := bson.M{"questions": bson.M{"$regex": primitive.Regex{Pattern: searchTerm, Options: "i"}}}
	filter := bson.M{
        "$text": bson.M{"$search": searchTerm},
    }
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

	return questions, nil
}

func GetQuestionsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is search term 0:")
	vars := mux.Vars(r) 
	pageStr := vars["page"]
	searchTerm := r.URL.Query().Get("q")
    fmt.Println("this is search term",searchTerm)
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}

	var questions []model.Question
	var errSearch error

	if searchTerm != "" {
		questions, errSearch = SearchQuestionsByPage(page, searchTerm)
	} else {
		questions, errSearch = GetQuestionsByPage(page)
	}

	if errSearch != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Convert questions to JSON 
	jsonData, err := json.Marshal(questions)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}