package mongo

import (
"gopkg.in/mgo.v2"
"gopkg.in/mgo.v2/bson"
	"ui-mockup-backend"
)

type standardModel struct {
	Controls[] root.Controls
	StandardName bson.ObjectId `bson:"_id,omitempty"`
}

func standardsIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"standardName"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

func newStandardModel(std *root.Standard) (*standardModel) {
	standard := standardModel{StandardName: std.StandardName, Controls:std.Controls}
	return &standard
}