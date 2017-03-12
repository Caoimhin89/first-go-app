package models

import "gopkg.in/mgo.v2/bson"

type Task struct {
	ID bson.ObjectId 'bson:"_id,omitempty" json:"id"'
	Name string 'bson:"name" json:"name"'
}

var Tasks = new(tasks)
type tasks struct{}