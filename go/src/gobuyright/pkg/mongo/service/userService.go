package service

import (
	"gobuyright/pkg/entity"
	"gobuyright/pkg/mongo"
	"gobuyright/pkg/mongo/model"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"
)

type UserService struct {
	collection *mgo.Collection
}

func NewUserService(session *mongo.Session, dbName string, colName string) *UserService {
	collection := session.GetCollection(dbName, colName)
	collection.EnsureIndex(model.UserModelIndex())
	return &UserService{collection}
}

func (us *UserService) CreateUser(u *entity.User) error {
	user := model.NewUserModel(u)
	return us.collection.Insert(&user)
}

func (us *UserService) GetByUsername(username string) (*entity.User, error) {
	m := model.UserModel{}
	err := us.collection.Find(bson.M{"username": username}).One(&m)
	return model.ToUser(&m), err
}
