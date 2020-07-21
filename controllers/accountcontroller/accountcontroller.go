package accountcontroller

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"html/template"
	"log"
	"net/http"
)
//var store =sessions.NewCookieStore([]byte("mysession"))
//var tpl=template.Must(template.ParseFiles("view/accountcontroller/index.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("<h2>thoufeek</h2>"))
	t,_:=template.ParseFiles("view/accountcontroller/index.html")
	t.Execute(w,nil)
	r.ParseForm()
	fmt.Println(r.Form)
	var username = r.FormValue("username")
	var password = r.FormValue("password")
	Authenticate(username,password)
}

func Authenticate(username, password string) {
	connection, err := mgo.Dial("localhost:27017")
	if err != nil {
		return
	}
	defer connection.Close()
	type credentials1 struct {
		Username   string     `bson:"username"`
		Password   string        `bson:"password"`
	}
	type credentials2 struct {
		Username   string     `bson:"username"`
		Password   string        `bson:"password"`
	}
	var detailsfromfront =[]credentials1{{username,password}}
	log.Println(detailsfromfront)
	data := connection.DB("Webpage").C("login")
	//data:=connection.Copy()
	fmt.Println("Connected to MongoDB!")
	for _, value := range detailsfromfront {
		in, _ := data.Upsert(value, value)
		if in != nil {
			fmt.Println(err)
		}
		detailsfromback:= []credentials2{}
		err:=data.Find(bson.M{}).All(&detailsfromback)
		if err !=nil{
			log.Println("detailsfromback",err)
		}
		if username ==value.Username && password == value.Password{
			fmt.Println("success")
		} else {
			fmt.Println("invalid")
		}
	}
}
/*	func Login(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Form)
	log.Println("path", r.URL.Path)
	fmt.Println(r.Form["url_long"])
	fmt.Println("method", r.Method)

	log.Printf("Login")
	temp, err := template.ParseFiles("view/accountcontroller/welcome.html")
	log.Printf("welcome", err)
	temp.Execute(w, nil)
}*/

/*func Welcome(w http.ResponseWriter, r *http.Request)  {
	log.Printf("Welcome")
	session,_:=store.Get(r,"mysession")
	username:=session.Values["username"]
	password:=session.Values["password"]
fmt.Println(username)
	fmt.Println(password)
	data :=map[string]interface{}{
		"username":username,
		"password":password,
	}

	temp,_:=template.ParseFiles("view/accountcontroller/welcome.html")
	temp.Execute(w,data)
}*/