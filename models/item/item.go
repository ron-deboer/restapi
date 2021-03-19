package models

import (
	conn "RESTapi/db"
	"gopkg.in/mgo.v2/bson"
)

// Item db schema
type Item struct {
	ID                bson.ObjectId `bson:"_id"`
	Title             string        `bson:"title"`
	Disporder         int           `bson:"disporder"`
	Boardcode         string        `bson:"boardcode"`
	Projectcode       string        `bson:"projectcode"`
	Prioritycode      string        `bson:"prioritycode"`
	Sizecode          string        `bson:"sizecode"`
	Createdbyuser     string        `bson:"createdbyuser"`
	Createdtimestamp  string        `bson:"createdtimestamp"`
	Assignedtouser    string        `bson:"assignedtouser"`
	Assignedtimestamp string        `bson:"assignedtimestamp"`
	Closedbyuser      string        `bson:"closedbyuser"`
	Closedtimestamp   string        `bson:"closedtimestamp"`
	Description       string        `bson:"description"`
	Comments          string        `bson:"comments"`
	Duedate           string        `bson:"duedate"`
}

type Items []Item

func ItemInfo(id bson.ObjectId, ItemCollection string) (Item, error) {
	db := conn.GetMongoDB()
	item := Item{}
	err := db.C(ItemCollection).Find(bson.M{"_id": &id}).One(&item)
	return item, err
}
