package model

import (
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Question struct {
	// QuestionId primitive.ObjectID `bson:"_id, omitempty"`
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Questions  string        `josn:"questions, omitempty"`
	Answer     []*Answer     `json:"answer"`
}
type Answer struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Answered string `bson:"answered"`
	AnsBy    string `bson:"ansby"`
	Upvote   int     `bson:"upvote"`
	UpvotedBy []string `bson:"uBy"`
}
type User struct{
	ID  primitive.ObjectID `bson:"_id,omitempty"`
	Email string  `bson:"email"`
	Name string   `bson:"name"`
}
