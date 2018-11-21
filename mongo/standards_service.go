package mongo

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"ui-mockup-backend"
)

type StandardsService struct {
	stdCollection *mgo.Collection
	certCollection *mgo.Collection
}

func NewStandardsService(session *mgo.Session, config *root.MongoConfig) *StandardsService {
	Stdcollection := session.DB(config.DbName).C("std")
	Certcollection := session.DB(config.DbName).C("cert")
	Stdcollection.EnsureIndex(standardsIndex())
	Certcollection.EnsureIndex(certificationIndex())
	return &StandardsService{stdCollection:Stdcollection, certCollection:Certcollection}
}

func (p *StandardsService) CreateStandard(std *root.Standard) error {
	standard := newStandardModel(std)
	return p.stdCollection.Insert(&standard)
}

func (p *StandardsService) CreateCertification(u *root.Certification) error {
	fmt.Println(u);
	return p.certCollection.Insert(&u)
}

func (p *StandardsService) GetStandardInfo(standardName string) (error, root.Standard) {
	standardsModel := standardModel{}
	err := p.stdCollection.Find(bson.M{"standardName": standardName}).One(&standardsModel)
	return err, root.Standard{
		StandardName: standardsModel.StandardName,
		Controls: standardsModel.Controls}
}
