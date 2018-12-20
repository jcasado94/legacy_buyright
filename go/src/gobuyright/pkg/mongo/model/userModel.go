package model

import (
	"gobuyright/pkg/entity"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserModel is the DB model for User
type UserModel struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Username string
}

// UserModelIndex constructs the mgo.Index for userModel
func UserModelIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

// NewUserModel creates a new UserModel given a User
func NewUserModel(u *entity.User) *UserModel {
	return &UserModel{
		Username: u.Username,
	}
}

func (um *UserModel) ToUser() *entity.User {
	return &entity.User{
		ID:       um.ID.Hex(),
		Username: um.Username,
	}
}
