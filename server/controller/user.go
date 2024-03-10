package mong

import (
	// "fmt"
	"context"
	"encoding/json"

	// "fmt"
	"io"
	"log"
	"net/http"

	"github.com/rishuraj2401/quest/model"
	"go.mongodb.org/mongo-driver/bson"
)

func HandleSignIn(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	var user model.User
	json.Unmarshal(data, &user)
	// Check if user with the given email exists
	existingUser := userCol.FindOne(context.TODO(), bson.M{"email": user.Email})

	// Check if there was an error during the FindOne operation
	if existingUser.Err() != nil {
		// log.Fatal("error",existingUser.Err())
		_, err := userCol.InsertOne(context.TODO(), user)
		if err != nil {
			log.Fatal(err)
		}
		w.Write([]byte("User inserted successfully"))
	}

	// If the user with the given email exists
	if existingUser.Err() == nil {
		w.Write([]byte("User already exists"))
	}
}
