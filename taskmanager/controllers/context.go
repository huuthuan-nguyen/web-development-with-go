package controllers

import (
	"gopkg.in/mgo.v2"
	"github.com/shijuvar/go-web/taskmanager/common"
)

// struct used for maintaining HTTP Request Context
type Context struct {
	MongoSession *mgo.Session
}

// close mgo.Session
func (c *Context) Close() {
	c.MongoSession.Close()
}

// Return mgo.collection for the given name
func (c *Context) DbCollection(name string) *mgo.Collection {
	return c.MongoSession.DB(common.AppConfig.Database).C(name)
}

// create a new context object for each http request
func NewContext() *Context {
	session := common.GetSession().Copy()
	context := &Context{
		MongoSession: session,
	}
	return context
}