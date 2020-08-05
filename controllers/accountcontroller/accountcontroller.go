package accountcontroller

import (
	"awesomeProject1/DB"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"html/template"
	"log"
	"net/http"

)

type User struct{
	UserId         bson.ObjectId `bson:"_id"`
	Username       string      `bson:"username"`
	Password       string      `bson:"password"`
}

func Index(w http.ResponseWriter, r *http.Request) {

	t,_:=template.ParseFiles("view/accountcontroller/index.html")
	t.Execute(w,nil)
	r.ParseForm()
	fmt.Println(r.Form)
	var username = r.FormValue("username")
	var password = r.FormValue("password")
	Authenticate(username,password)
}

func Authenticate(username, password string) {

	type credentials1 struct {
		Username   string     `bson:"username"`
		Password   string        `bson:"password"`
	}
	type credentials2 struct {
		Username   string     `bson:"username"`
		Password   string        `bson:"password"`
	}
	dbsession := DB.Obj.Copy()
	defer dbsession.Close()

	var detailsfromfront =[]credentials1{{username,password}}
	log.Println(detailsfromfront)
	data := dbsession.DB("Webpage").C("login")
	//data:=connection.Copy()
	fmt.Println("Connected to MongoDB!")
	for _, value := range detailsfromfront {
		in, err := data.Upsert(value, value)
		if in != nil {
			fmt.Println(err)
		}
		detailsfromback:= []credentials2{}
		err = data.Find(bson.M{}).All(&detailsfromback)
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
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	t,_:=template.ParseFiles("view/accountcontroller/createAccount.html")
	t.Execute(w,nil)

	r.ParseForm()
	fmt.Println(r.Form)
	log.Println("CreateAccount Page")

	session:=DB.Obj.Copy()
	defer session.Close()

	type Signup struct {
		Username    string `json:"username"`
		Password    string `json:"password"`
	} 
	res :=make(map[string]interface{})
	reg :=Signup{}
	err:=session.DB("Webpage").C("login").Insert(reg)
	if err !=nil{
		log.Println("err",err)
	}
	var username = r.FormValue("username")
	log.Println("username",username)
	var password = r.FormValue("password")
	log.Println("password",password)
	res["name"]=Signup{Username: username,Password: password}
	res["message"]="MESSAGE"

	
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	t,_:=template.ParseFiles("view/accountcontroller/deleteAccount.html")
	t.Execute(w,nil)
	r.ParseForm()
	fmt.Println(r.Form)
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