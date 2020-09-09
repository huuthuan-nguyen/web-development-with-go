package common

import (
	"gopkg.in/mgo.v2"
	"time"
)

var session *mgo.Session

func GetSession() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs: []string{AppConfig.MongoDBHost},
			Username: AppConfig.DBUser,
			Password: AppConfig.DBPwd,
			Timeout: 60 * time.Second,
		})

		if err != nil {
			log.Fatalf("[GetSession]: %s\n", err)
		}
	}

	return session
}

func createDBSession() {
	var err error

	session, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs: []string{AppConfig.MongoDBHost},
		Username: AppConfig.DBUser,
		Password: AppConfig.DBPwd,
		Timeout: 60 * time.Second,
	})

	if err != nil {
		log.Fatalf("[CreateDBSession]: %s\n", err)
	}
}

// Add Indexes to MongoDB
func addIndexes() {
	var err error
	userIndex = mgo.Index{
		Key: []string{"email"},
		Unique: true,
		Background: true,
		Sparse: true,
	}

	taskIndex = mgo.Index{
		Key: []string{"createdby"},
		Unique: false,
		Background: true,
		Sparse: true,
	}

	noteIndex = mgo.Index{
		Key: []string{"taskid"},
		Unique: false,
		Background: true,
		Sparse: true,
	}

	// Add Indexes into MongoDB
	session := GetSession().Copy()
	defer session.Close()

	userCol := session.DB(AppConfig.Database).C("users")
	taskCol := session.DB(AppConfig.Database).C("tasks")
	noteCol := session.DB(AppConfig.Database).C("notes")

	err = session.EnsureIndex(userIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}
	err = session.EnsureIndex(taskIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}
	err = session.EnsureIndex(noteIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}
}