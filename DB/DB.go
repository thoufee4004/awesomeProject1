package DB

import (
	"github.com/globalsign/mgo"
	"log"
	"sync"
)
var(
	Obj    *mgo.Session
	once    sync.Once
)

/*func GetDB() {
	once.Do(func() {
		Obj = DBconnect()
	})
}*/
// DB connection
func DBconnect() (*mgo.Collection, error) {
	//client, err := NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	mgosession, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatal("cannot dial mongo", err)
	}
	//fmt.Println("db connected..")
	//return mgosession
//}
	err = mgosession.Ping()
	if err != nil {
		return nil, err
	}
	collection := mgosession.DB("Webpage").C("login")
	return collection, nil
}

func GetDb() *mgo.Database{
	dbObj := Obj.Copy()
	return dbObj.DB("Webpage")
}
