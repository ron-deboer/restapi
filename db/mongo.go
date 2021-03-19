package conn

import (
	"fmt"
	"os"

	mgo "gopkg.in/mgo.v2"
	"github.com/joho/godotenv"
)

var db *mgo.Database

func init() {
	err := godotenv.Load(".env")
	host := os.Getenv("MONGO_HOST")
	dbName := os.Getenv("MONGO_DB_NAME")
	fmt.Printf("mongo host: [%s]  db: [%s]\n", host, dbName)
	session, err := mgo.Dial(host)
	if err != nil {
		fmt.Println("mongo session err:", err)
		os.Exit(2)
	}
	db = session.DB(dbName)
}

// return DB connection
func GetMongoDB() *mgo.Database {
	return db
}
