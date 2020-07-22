package main

import (
	"awesomeProject1/controllers/accountcontroller"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {

	//Port
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	var Router = mux.NewRouter()
	var S = Router.PathPrefix("/api").Subrouter()

	//CSS file
	FileServerObj := http.FileServer(http.Dir("assets"))
	S.Handle("/assets/", http.StripPrefix("/assets/", FileServerObj))

	//Routes

	S.HandleFunc("/createUserAccount", accountcontroller.CreateAccount).Methods("POST")

	S.HandleFunc("/getUserAccount", accountcontroller.Index).Methods("POST")

	S.HandleFunc("/deleteUserAccount{id}", accountcontroller.DeleteAccount).Methods("DELETE")

	//db
	DBconnect()

	//Listening server
	err := http.ListenAndServe(""+port, Router) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: Not Listening..", err)
	}

}

// DB connection
func DBconnect() *mgo.Session {
	dbsession := Obj.Copy()
	defer dbsession.Close()
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatal("cannot dial mongo", err)
		os.Exit(1)
	}
	fmt.Println("db connected..")
	return session
}


/*package main

import (
"fmt"
"github.com/globalsign/mgo"
"github.com/globalsign/mgo/bson"
"log"
)

type Student struct {
	Name   string     `bson:"name"`
	Age    int        `bson:"age"`
	Mark   int        `bson:"mark"`

}

func main() {
	st:=[]Student{
		{"tfk", 22, 95},
		{"iron man",45,66},
		{"edith",2,100},
		{"tony",45,100},
		{"thor",55,90},
		{"hulk",55,98},
		{"cap",60,88},
	}

	session,err:=mgo.Dial("localhost:27017")
	if err != nil {
		return
	}
	con:=session.DB("Task").C("students")

	len, err := con.Count()
	if err != nil {
		return
	}
	log.Println("len :",len)

	fmt.Println("Connected to MongoDB!")
	for _,record:= range st {
		in, _ :=con.Upsert(record,record)
		if in != nil {
			fmt.Println(err)
		}
	}
	var res []Student
	err = con.Find(bson.M{}).All(&res)
	if err != nil {
		return
	}
	//fmt.Println(res)
	fmt.Println("inserted: ", res)

}
//Rverse
func Reverse{
	for i,j:=0,len(a)-1;i < j ;i,j=i+1,j-1{
		a[i],a[j]=a[j],a[i]
	}
}

//input [1,2,3,4]
//output [1,3,6,10]
/*func main()  {
	a:=[]int{1,2,3,4,5}
	b:=0
	for i,_:= range a{
		b+=a[i]
		fmt.Println("output",b)
	}
}
*/
/*type Page struct {
	Title string
	Body string
}

func PageHandler(w http.ResponseWriter, r *http.Request)  {
	a:= Page{Title: "Thoufeek", Body: "this my first page ever"}
	t,err:=template.ParseFiles("createAccount.html")
	fmt.Println(err)
	t.Execute(w,a)

}
func handlerFunc(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w,"<h1>This is my page Thoufeek attempt1</h1>")

}*/