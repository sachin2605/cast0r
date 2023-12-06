package svc

import (
	"github.com/sachin2605/cast0r/db"
	"go.mongodb.org/mongo-driver/mongo"
)

var DB *mongo.Database

func InitServices() {
	DB = db.ConnectDB().Database("Fruits")
	return
}
