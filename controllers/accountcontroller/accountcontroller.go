package accountcontroller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"pkg/mod/github.com/gorilla/sessions"
)
var store =sessions.NewCookieStore([]byte("mysession"))

func Index(w http.ResponseWriter, r *http.Request){
	temp,_:=template.ParseFiles("view/accountcontroller/index.html")
	log.Printf("index",r.URL.Path,r.Form)
	temp.Execute(w,nil)
	http.Redirect(w,r,"view/accountcontroller/welcome.html",http.StatusSeeOther)
}

func Login(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	fmt.Println("method",r.Method)
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	fmt.Println("username",username)
	fmt.Println(password)

	log.Printf("Login")
	temp,err:=template.ParseFiles("view/accountcontroller/welcome.html")
	log.Printf("welcome",err)
	temp.Execute(w,nil)

	if username=="asd" && password =="123" {
		log.Printf("username",username)
		session,_:=store.Get(r,"mysession")
		session.Values["username"]= username
		session.Values["password"]= password
		fmt.Println("user",username)

	} else {
		log.Printf("invalid")
		data :=map[string]interface{}{
			"err":"invalid",
		}
		temp, _ := template.ParseFiles("view/accountcontroller/index.html")
		//http.Redirect(w,r,"view/accountcontroller/index",http.StatusSeeOther)
		temp.Execute(w, data)


	}
}

func Welcome(w http.ResponseWriter, r *http.Request)  {
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
}