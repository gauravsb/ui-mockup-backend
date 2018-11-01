package mongo

import (
	"gopkg.in/mgo.v2"
)

type Standards struct {
	ControlName string
	ControlInfo ControlInfo
}

type ControlInfo struct {
	Family          string  `json:"family"`
	Name     		string  `json:"name"`
	Description     string  `json:"desc"`
}

type Standards_Model struct {
	Standards Standards
	StandardName string
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
