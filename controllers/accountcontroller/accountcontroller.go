package accountcontroller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"pkg/mod/github.com/gorilla/sessions"
)
var store =sessions.NewCookieStore([]byte("mysession"))
var tpl=template.Must(template.ParseFiles("view/accountcontroller/index.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>thoufeek</h1>"))
	tpl.Execute(w ,nil)
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println("username", username)
	fmt.Println("password", password)

/*	if username == "asd" && password == "123" {
		log.Printf("username", username)
		session, _ := store.Get(r, "mysession")
		session.Values["username"] = username
		session.Values["password"] = password
		fmt.Println("user", username)

	} else {
		log.Printf("invalid")
		data := map[string]interface{}{
			"err": "invalid",
		}
		temp, _ := template.ParseFiles("view/accountcontroller/index.html")
		//http.Redirect(w,r,"view/accountcontroller/index",http.StatusSeeOther)
		temp.Execute(w, data)
	}*/
}
func Login(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Form)
	log.Println("path", r.URL.Path)
	fmt.Println(r.Form["url_long"])
	fmt.Println("method", r.Method)

	log.Printf("Login")
	temp, err := template.ParseFiles("view/accountcontroller/welcome.html")
	log.Printf("welcome", err)
	temp.Execute(w, nil)
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