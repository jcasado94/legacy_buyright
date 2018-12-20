package model

import (
	"gobuyright/pkg/entity"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UsageModel struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	UsageID   string
	UsageName string
}

func UsageModelIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"usageID"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

func NewUsageModel(u *entity.Usage) *UsageModel {
	return &UsageModel{
		UsageID:   u.UsageID,
		UsageName: u.UsageName,
	}
}

func (um *UsageModel) ToUsage() *entity.Usage {
	return &entity.Usage{
		ID:        um.ID.Hex(),
		UsageID:   um.UsageID,
		UsageName: um.UsageName,
	}
}
