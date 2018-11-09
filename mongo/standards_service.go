package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"ui-mockup-backend"
)

type StandardsService struct {
	collection *mgo.Collection
}

func NewStandardsService(session *mgo.Session, config *root.MongoConfig) *StandardsService {
	collection := session.DB(config.DbName).C("std")
	collection.EnsureIndex(standardsIndex())
	return &StandardsService{collection}
}

func (p *StandardsService) CreateStandard(std *root.Standard) error {
	standard := newStandardModel(std)
	return p.collection.Insert(&standard)
}

func (p *StandardsService) GetStandardsInfo(standardName string) (error, root.Standard) {
	standardsModel := standardModel{}
	err := p.collection.Find(bson.M{"standardName": standardName}).One(&standardsModel)
	return err, root.Standard{
		StandardName: standardsModel.StandardName,
		Controls: standardsModel.Controls}
}
