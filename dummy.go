













///////package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)
type USER struct {
	Username string      `bson:"Username" json:"Username,omitempty"`
	Password string      `bson:"Password" json:"Password,omitempty"`
}
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!")
}
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, nil)
	} else {
		//connection with mongodb
		session, err := mgo.Dial("mongodb://127.0.0.1:27017/so")
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB("so").C("insrt")
		doc := USER{
			Username: r.Form["username"][0],
			Password: r.Form["password"][0],
		}
		err = c.Insert(doc)
		if err != nil {
			panic(err)
		}
		r.ParseForm()
		fmt.Printf("%T\n", r.Form["username"])
		fmt.Println("---------------------------------------------")
		fmt.Println("username:", r.Form["username"][0])//output:- username: user_name
		fmt.Println("password:", r.Form["password"][0])//password:- password: user_password
	}
}
func main() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
