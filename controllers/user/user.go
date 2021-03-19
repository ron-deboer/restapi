package user

import (
	"errors"
	"net/http"

	"RESTapi/conn"
	user "RESTapi/models/user"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// Users Collection 
const UserCollection = "users"

var (
	errNoData          = errors.New("No matching data found")
)

// Get All users
func GetAll(c *gin.Context) {

	db := conn.GetMongoDB()

	users := user.Users{}
	err := db.C(UserCollection).Find(bson.M{}).All(&users)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": errNoData.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "users": &users})
}

// Get User 
func Get(c *gin.Context) {
	var id bson.ObjectId = bson.ObjectIdHex(c.Param("id")) 
	user, err := user.UserInfo(id, UserCollection)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": errNoData.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "user": &user})
}

// Create User
func Create(c *gin.Context) {

	db := conn.GetMongoDB()

	user := user.User{}
	err := c.Bind(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": errNoData.Error()})
		return
	}
	user.ID = bson.NewObjectId()


	err = db.C(UserCollection).Insert(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": errNoData.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "user": &user})
}

// Update User
func Update(c *gin.Context) {

	db := conn.GetMongoDB()

	var id bson.ObjectId = bson.ObjectIdHex(c.Param("id")) 
	existingUser, err := user.UserInfo(id, UserCollection)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": errNoData.Error()})
		return
	}
	err = c.Bind(&existingUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": errNoData.Error()})
		return
	}

	existingUser.ID = id

	err = db.C(UserCollection).Update(bson.M{"_id": &id}, existingUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": errNoData.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "user": &existingUser})
}

// Delete User
func Delete(c *gin.Context) {

	db := conn.GetMongoDB()

	var id bson.ObjectId = bson.ObjectIdHex(c.Param("id"))

	err := db.C(UserCollection).Remove(bson.M{"_id": &id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": errNoData.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User deleted"})
}
