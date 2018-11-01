package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type StandardsService struct {
	collection *mgo.Collection
}

func NewStandardsService(session *mgo.Session, config *root.MongoConfig) *StandardsService {
	collection := session.DB(config.DbName).C("std")
	collection.EnsureIndex(standardsIndex())
	return &StandardsService{collection}
}

func(p *StandardsService) CreateStandard(std *root.Standards) error {
	return p.collection.Insert(&std)
}

func (p *StandardsService) GetStandardsInfo(standardName string) (error, Standards) {
	standardsModel := Standards{}
	err := p.collection.Find(bson.M{"StandardName": standardName}).One(&standardsModel)
	return err, standardsModel}
}

