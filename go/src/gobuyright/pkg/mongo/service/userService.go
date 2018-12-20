package service

import (
	"gobuyright/pkg/entity"
	"gobuyright/pkg/mongo"
	"gobuyright/pkg/mongo/model"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"
)

// UserService serves the mongo operations for User.
type UserService struct {
	collection *mgo.Collection
}

// NewUserService creates a new UserService given database and collection names, connection through session.
func NewUserService(session *mongo.Session, dbName string, colName string) *UserService {
	collection := session.GetCollection(dbName, colName)
	collection.EnsureIndex(model.UserModelIndex())
	return &UserService{collection}
}

// CreateUser inserts u into the collection.
func (us *UserService) CreateUser(u *entity.User) error {
	user := model.NewUserModel(u)
	return us.collection.Insert(&user)
}

// GetByUsername retrieves the User with Username username from the collection.
func (us *UserService) GetByUsername(username string) (*entity.User, error) {
	um := model.UserModel{}
	err := us.collection.Find(bson.M{"username": username}).One(&um)
	return um.ToUser(), err
}
