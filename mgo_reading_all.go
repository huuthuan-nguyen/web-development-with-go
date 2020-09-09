package main

import (
	"time"
	"fmt"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Task struct {
	Description string
	Due time.Time
}

type Category struct {
	Id bson.ObjectId `bson:"_id,omitempty"`
	Name string
	Description string
	Tasks []Task
}

func main() {

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	// get collection
	c := session.DB("taskdb").C("categories")

	// query the "categories" collection
	iter := c.Find(nil).Sort("-name").Iter()

	result := Category{}
	for iter.Next(&result) {
		fmt.Printf("Category:%s, Description:%s\n", result.Name, result.Description)
		tasks := result.Tasks
		for _, v := range tasks {
			fmt.Printf("Task:%s Due:%v\n", v.Description, v.Due)
		}
	}
	if err = iter.Close(); err != nil {
		log.Fatal(err)
	}
}