package mongo

import (
	"gobuyright/pkg/entity"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"
)

type GfUserService struct {
	collection *mgo.Collection
}

func NewGfUserService(session *Session, dbName string, colName string) *GfUserService {
	collection := session.GetCollection(dbName, colName)
	collection.EnsureIndex(gfUserModelIndex())
	return &GfUserService{collection}
}

func (s *GfUserService) CreateUser(u *entity.GfUser) error {
	user := newGfUserModel(u)
	return s.collection.Insert(&user)
}

func (s *GfUserService) GetByUsername(username string) (*entity.GfUser, error) {
	model := gfUserModel{}
	err := s.collection.Find(bson.M{"username": username}).One(&model)
	return model.toGfUser(), err
}
