package DB

import (
	"fmt"
	"github.com/globalsign/mgo"
	"log"
	"sync"
)
var(
	Obj    *mgo.Session
	once    sync.Once
)

func GetDB() {
	once.Do(func() {
		Obj = DBconnect()
	})
}
// DB connection
func DBconnect() *mgo.Session {
	Session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatal("cannot dial mongo", err)
	}
	fmt.Println("db connected..")
	return Session
}
/*func GetDb() *mgo.Database{
	dbObj := Obj.Copy()
	return dbObj.DB("Webpage")
}*/
