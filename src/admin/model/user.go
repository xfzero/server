package model

import (
	"gm/base"

	"go.mongodb.org/mongo-driver/bson"
)

var cName = "users"

func GetUserLogin(username, password string) *base.AdminUser {
	user := &base.AdminUser{}

	not, _ := FindOne(cName, bson.M{"username": username}, user)
	if not {
		return nil
	}

	return user
}
