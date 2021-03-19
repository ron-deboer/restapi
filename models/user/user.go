package model

import (
	"RESTapi/conn"
	"gopkg.in/mgo.v2/bson"
)

// User db schema
type User struct {
	ID        		bson.ObjectId `bson:"_id"`
	Username      	string        `bson:"username"`
	Name   			string        `bson:"name"`
	Email       	string        `bson:"email"`
	Password 		string     	  `bson:"password"`
	Boardcode 		string        `bson:"boardcode"`
}

type Users []User

func UserInfo(id bson.ObjectId, userCollection string) (User, error) {
	db := conn.GetMongoDB()
	user := User{}
	err := db.C(userCollection).Find(bson.M{"_id": &id}).One(&user)
	return user, err
}
