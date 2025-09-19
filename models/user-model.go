package models

import (
	"strings"

	"go.mongodb.org/mongo-driver/v2/bson"
)


type UserRole string

const (
	User		UserRole = "User"
	Admin		UserRole = "Admin"
);

// Optional: Add a String() method for easy printing
func (s *UserRole) String() string {
	return string(*s)
}

// Map for string representation and reverse lookup
var roleNames = map[UserRole]string{
	User:   "User",
	Admin: 	"Admin",
}

// Parse function for string-to-enum
func GetRole(name string) (UserRole, bool) {
	for key, val := range roleNames {
		if strings.EqualFold(val, name) {
			return key, true
		}
	}

	return User, false
}

type UserModel struct {
	ID          bson.ObjectID 		`json:"id" bson:"_id,omitempty"`
	FirstName		string 		`json:"first_name" binding:"required" validate:"gte=1" bson:"first_name"`
	LastName		string 		`json:"last_name" binding:"required" validate:"gte=1" bson:"last_name"`
	Email				string 		`json:"email" binding:"required" validate:"email" bson:"email"`
	Role				UserRole 	`json:"role" binding:"required"`
	Articles		[]ArticleModel `json:"articles" binding:"required"`
}

// func (u *UserModel) GetBirthday() string {
// 	return u.DoB.Format("yyyyMMddHHmmss")
// }
