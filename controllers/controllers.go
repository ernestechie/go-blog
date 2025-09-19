package controllers

import (
	"github.com/ernestechie/go-blog/db"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var usersColl *mongo.Collection

// Initialize mongodb client
func init()  {
	c := db.ConnectDB()
	usersColl = c.Database("db").Collection("users")
}
