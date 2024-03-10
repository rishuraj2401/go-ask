package mong 

import (
	"context"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/rishuraj2401/quest/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func upvote(qId string, aId string, userId string) {
	id, _ := primitive.ObjectIDFromHex(qId)
	aID, _ := primitive.ObjectIDFromHex(aId)
	filter := bson.M{
		"_id": id, "answer": bson.M{"$elemMatch": bson.M{"_id": aID}},
	}

	// var q1 model.Question
	filter1 := bson.M{
		"_id":        id,
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
	arr := question.Answer[0].UpvotedBy
	add := 0
	if len(arr) == 0 {
		add++
		arr = append(arr, userId)
	} else {
		var bool1 bool = true
		for ind, val := range arr {
			if val == userId {
				add--
				arr = append(arr[:ind], arr[ind+1:]...)
				bool1 = false
				break
			}
		}
		if bool1 {
			add++
			arr = append(arr, userId)
		}
	}
	fmt.Println(add, "this is the arr")
	// 'questions' now contains the matching Question documents with only the specified answer

	fmt.Println("something is here in upvote")
	update := bson.M{
		"$set": bson.M{
			"answer.$.upvote": question.Answer[0].Upvote + add,
			"answer.$.uBy":    arr,
		},
		
	}
	// fmt.Println(update, "ans:", ans.Upvote)
	result, err := collection.UpdateOne(context.Background(), filter, update)
	fmt.Println(result, "err", err)
	// return question

}


func Like(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is router 4")
	w.Header().Set("content-type", "application/json")
	param := mux.Vars(r)
	aId := param["aId"]
	qId := param["qId"]
	user := param["user"]
	// // data,_:=io.ReadAll(r.Body)
	// var ans model.Answer
	// err:= json.Unmarshal(data,&ans)
	// if err!=nil{
	// 	log.Fatal(err)
	// }
	fmt.Println("answer is updated", aId, qId)
	upvote(qId, aId, user)
	// jsonD,_:= json.Marshal(res)
	// fmt.Println("answer is updated",param["id"])
	w.Write([]byte("like is working"))
}