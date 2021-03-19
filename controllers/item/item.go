package controllers

import (
	"errors"
	"net/http"

	conn "RESTapi/db"
	item "RESTapi/models/item"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// Items Collection 
const ItemCollection = "items"

var (
	errNoData          = errors.New("No matching data found")
)

// Get All items
func GetAll(c *gin.Context) {

	db := conn.GetMongoDB()

	items := item.Items{}
	err := db.C(ItemCollection).Find(bson.M{}).All(&items)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": errNoData.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "items": &items})
}

// Get Item 
func Get(c *gin.Context) {
	var id bson.ObjectId = bson.ObjectIdHex(c.Param("id")) 
	item, err := item.ItemInfo(id, ItemCollection)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": errNoData.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "item": &item})
}

// Create Item
func Create(c *gin.Context) {

	db := conn.GetMongoDB()

	item := item.Item{}
	err := c.Bind(&item)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": errNoData.Error()})
		return
	}
	item.ID = bson.NewObjectId()


	err = db.C(ItemCollection).Insert(item)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": errNoData.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "item": &item})
}

// Update Item
func Update(c *gin.Context) {

	db := conn.GetMongoDB()

	var id bson.ObjectId = bson.ObjectIdHex(c.Param("id")) 
	existingItem, err := item.ItemInfo(id, ItemCollection)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": errNoData.Error()})
		return
	}
	err = c.Bind(&existingItem)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": errNoData.Error()})
		return
	}

	existingItem.ID = id

	err = db.C(ItemCollection).Update(bson.M{"_id": &id}, existingItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": errNoData.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "item": &existingItem})
}

// Delete Item
func Delete(c *gin.Context) {

	db := conn.GetMongoDB()

	var id bson.ObjectId = bson.ObjectIdHex(c.Param("id"))

	err := db.C(ItemCollection).Remove(bson.M{"_id": &id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": errNoData.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Item deleted"})
}
