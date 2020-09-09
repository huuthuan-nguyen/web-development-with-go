package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type (
	User struct {
		Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
		FirstName string `json:"firstname"`
		LastName string `json:"lastname"`
		Email string `json:"email"`
		Password string `json:"password,omitempty"`
		HashPassword []byte `json:"hashpassword,omitempty"`
	}
	Task struct {
		Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
		CreatedBy string `json:"createdby"`
		Description string `json:"description"`
		CreatedOn time.Time `json:"createdon,omitempty"`
		Status string `json:"status,omitempty"`
		Tag []string `json:"tags,omitempty"`
	}
	TaskNote struct {
		Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
		TaskId bson.ObjectId `json:"taskid"`
		Description string `json:"description"`
		CreatedOn time.Time `json:"createdon,omitempty"`
	}
)