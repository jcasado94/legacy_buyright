package mongo

import (
	"../../pkg/"
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

func (s *GfUserService) CreateUser(u *root.GfUser) error {
	user := newGfUserModel(u)
	return s.collection.Insert(&user)
}

func (s *GfUserService) GetByID(id string) (*root.GfUser, error) {
	model := gfUserModel{}
	err := s.collection.FindId(id).One(&model)
	return model.toGfUser(), err
}
