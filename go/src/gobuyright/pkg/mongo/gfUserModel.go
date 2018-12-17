package mongo

import (
	"../../pkg"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type gfUserModel struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Username string
}

func gfUserModelIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

func newGfUserModel(u *root.GfUser) *gfUserModel {
	return &gfUserModel{
		Username: u.Username,
	}
}

func (u *gfUserModel) toGfUser() *root.GfUser {
	return &root.GfUser{
		ID:       u.ID.Hex(),
		Username: u.Username,
	}
}
