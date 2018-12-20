package service

import (
	"gobuyright/pkg/entity"
	"gobuyright/pkg/mongo"
	"gobuyright/pkg/mongo/model"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UsageService struct {
	collection *mgo.Collection
}

func NewUsageService(session *mongo.Session, dbName string, colName string) *UsageService {
	collection := session.GetCollection(dbName, colName)
	collection.EnsureIndex(model.UsageModelIndex())
	return &UsageService{collection}
}

func (us *UsageService) CreateUsage(u *entity.Usage) error {
	usage := model.NewUsageModel(u)
	return us.collection.Insert(&usage)
}

func (us *UsageService) GetByUsageID(usageID string) (*entity.Usage, error) {
	um := model.UsageModel{}
	err := us.collection.Find(bson.M{"usageID": usageID}).One(&um)
	return um.ToUsage(), err
}
