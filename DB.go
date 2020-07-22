package main

import (
	"github.com/globalsign/mgo"

)
var Obj *mgo.Session

func GetDB() *mgo.Database{
	dbObj := Obj.Copy()
	return dbObj.DB("Webpage")
}
