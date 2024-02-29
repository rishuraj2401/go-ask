package mong

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rishuraj2401/quest/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
const dbName string= "dbQ"
const col string= "Questions"
var collection *mongo.Collection
func init() {
	if err:= godotenv.Load(); err!=nil {
		log.Fatal("error in loading env");
	}
	url:=os.Getenv("URL")
	// const url string = "mongodb+srv://rishuraj2401:Rishu%402002@cluster0.twrql.mongodb.net"
	clientOpt := options.Client().ApplyURI(url)
	client, err:=mongo.Connect(context.TODO(),clientOpt)
	if err!=nil{
		log.Fatal(err)
	}
	index := mongo.IndexModel{
		Keys: bson.M{
			"questions": "text",
		},
	}
	// _, err = client.Database(dbName).Collection(col).Indexes().CreateOne(context.TODO(), index)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println("Db is connected")
	 collection=client.Database(dbName).Collection(col)
	 collection.Indexes().CreateOne(context.TODO(),index)
	 fmt.Println("Coleection instance is ready",collection)
	 // Add this code after creating the collection



}

func insertUser(user model.User){
	data, err:=collection.InsertOne(context.Background(), user)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Inserted successfully", data)
}
func User(w http.ResponseWriter , r * http.Request){
	fmt.Println("this is router 1")
	w.Header().Set("content-type", "application/json");
	data,_ := io.ReadAll(r.Body)
	var user model.User
	err:=json.Unmarshal(data, &user)
	if err!=nil{
		log.Fatal(err)
	}
	insertUser(user)
    w.Write([]byte("user Inseted succesfully{user}",))
}
func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
 }

func InsertOne(quest model.Question){
	data, err:=collection.InsertOne(context.Background(), quest)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Inserted successfully", data)
}
func Inserted(w http.ResponseWriter , r * http.Request){
	fmt.Println("this is router 5")
	setupCorsResponse(&w, r)
	w.Header().Set("content-type", "application/json");
	
	data,_ := io.ReadAll(r.Body)
	
	var question model.Question
	if question.ID.IsZero(){
		question.ID= primitive.NewObjectID()
	}
	err:=json.Unmarshal(data, &question)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("this is body",question, string(data))
	InsertOne(question)
    w.Write([]byte("Inseted succesfully"))
}
func Hello(w http.ResponseWriter , r * http.Request){
	fmt.Println("this is router 0")
	w.Header().Set("content-type", "application/json");
    w.Write([]byte("Welcome succesfully"))
}
func update(qId string, ans model.Answer , user string){
    id,_ :=primitive.ObjectIDFromHex(qId)
	ans.AnsBy=user
	filter:=bson.M{"_id":id}
	fmt.Println("something is here",id)
	update:=bson.M{"$push":bson.M{
		"answer":ans}}
	uData,_:=collection.UpdateOne(context.Background(), filter, update)
	fmt.Println("answer is here",uData)

}
func AddAnswer(w http.ResponseWriter, r * http.Request){
fmt.Println("This is router 3")
w.Header().Set("content-type", "application/json")
param:=mux.Vars(r)
user:= param["user"]
data,_:=io.ReadAll(r.Body)
var ans model.Answer
err:= json.Unmarshal(data,&ans)
if ans.ID.IsZero(){
	ans.ID= primitive.NewObjectID()
}
if err!=nil{
	log.Fatal(err)
}

update(param["id"],ans, user)
w.Write([]byte("hey over"))
fmt.Println("answer is updated",param["id"])
}
func upvote(qId string, aId string,  userId string){
    id,_ :=primitive.ObjectIDFromHex(qId)
	aID,_:=primitive.ObjectIDFromHex(aId)
	filter := bson.M{
		"_id": id , "answer": bson.M{"$elemMatch": bson.M{"_id": aID}},
		
	} 
	
    // var q1 model.Question
	filter1 := bson.M{
		"_id": id,
		"answer._id": aID,
	}
	
	projection := bson.M{
		"answer.$": 1,
	}
	
	cursor, err := collection.Find(context.TODO(), filter1, options.Find().SetProjection(projection))
	if err != nil {
		// Handle error
		panic(err)
	}
	
	defer cursor.Close(context.TODO())
	
	// var questions [] model.Question
	var question model.Question
	for cursor.Next(context.TODO()) {
		
		err := cursor.Decode(&question)
		if err != nil {
			// Handle error
		panic(err)

		}
		// questions = append(questions, question)
	}
	arr:= question.Answer[0].UpvotedBy
	add:=0
	if len(arr)==0{
		add++;
		arr=append(arr,userId )
	}else{
		var bool1 bool = true
	for ind,val :=range arr{
		if(val==userId){
          add--
		 arr= append(arr[:ind], arr[ind+1:]...)
		 bool1 =false
		 break;
		}
	} 
	if bool1{
	add++;
		arr=append(arr,userId )
	}}
	fmt.Println(add,"this is the arr")
	// 'questions' now contains the matching Question documents with only the specified answer
	
	fmt.Println("something is here in upvote")
	update := bson.M{
		"$set": bson.M{ 
			"answer.$.upvote": question.Answer[0].Upvote+add,
			"answer.$.uBy":arr,
		},
		// "$push":bson.M{
		// 	"answer.$.uBy":userId,
		// },
	}
	// fmt.Println(update, "ans:", ans.Upvote)
	result,err:=collection.UpdateOne(context.Background(), filter, update)
	fmt.Println(result,"err",err)
	// return question
	
}
func Like(w http.ResponseWriter, r * http.Request){
fmt.Println("This is router 4")
w.Header().Set("content-type", "application/json")
param:=mux.Vars(r)
aId:=param["aId"]
qId:=param["qId"]
user:=param["user"]
// // data,_:=io.ReadAll(r.Body)
// var ans model.Answer
// err:= json.Unmarshal(data,&ans)
// if err!=nil{
// 	log.Fatal(err)
// }
fmt.Println("answer is updated",aId, qId)
upvote(qId,aId,user)
// jsonD,_:= json.Marshal(res)
// fmt.Println("answer is updated",param["id"])
w.Write([]byte("like is working"))
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
func GetQuestion(w http.ResponseWriter, r *http.Request){
	param:= mux.Vars(r)
	page:=param["page"]
	pageNum,_:=strconv.Atoi(page)
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