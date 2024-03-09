// go get -u github.com/gorilla/mux
// go get go.mongodb.org/mongo-driver/mongo

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rishuraj2401/quest/controller"
	// "github.com/rishuraj2401/quest/model"
	"github.com/rs/cors" 
)

func main() {
	fmt.Println("hey it is running",)
	r:=mux.NewRouter()
	 
	// q1.Questions="What is question"
	r.HandleFunc("/",mong.Hello).Methods("GET")
	r.HandleFunc("/getQ/{page}",mong.GetQuestion).Methods("GET")
    r.HandleFunc("/search/{page}", mong.GetQuestionsHandler).Methods("GET")
    r.HandleFunc("/insert",mong.Inserted).Methods("POST")
    r.HandleFunc("/ans/{id}/{user}",mong.AddAnswer).Methods("PUT")
    r.HandleFunc("/like/{qId}/{aId}/{user}",mong.Like).Methods("GET")
    r.HandleFunc("/user",mong.Inserted).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "DELETE", "POST", "PUT"},
	})

	handler := c.Handler(r)
	log.Fatal((http.ListenAndServe(":8089", handler)))
	http.Handle("/", r)

	// mong.InsertOne(q1)
	// log.Fatal(http.ListenAndServe(":5000",r))
}
// var _ = Handler