package models

import "go.mongodb.org/mongo-driver/v2/bson"

type ArticleModel struct {
	ID           	bson.ObjectID 		`json:"id" bson:"_id,omitempty"`
	Title					string `json:"title" binding:"required" validate:"gte=1" bson:"title,omitempty"`
	Description		string `json:"description,omitempty" binding:"required" validate:"gte=1" bson:"description,omitempty"`
	// Author				UserModel `json:"author" binding:"required" validate:"required" bson:"author,omitempty"`
	// Author				UserModel `json:"author"`
	Author 					bson.ObjectID		`json:"author"`
}
