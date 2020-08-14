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

	Username       string      `bson:"username"`
	Password       string      `bson:"password"`
	confirmpassword string     `bson:"confirmpassword"`
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
func Welcome(w http.ResponseWriter, r *http.Request){

	fmt.Println("method", r.Method)
	t, err := template.ParseFiles("view/accountcontroller/welcome.html")
	if err != nil {
		log.Println("error", err)
	}
	t.Execute(w, nil)
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("view/accountcontroller/createAccount.html")
	if err != nil {
		log.Println("parse error", err)
	}

	Name := r.FormValue("username") // Data from the form
	pwd := r.FormValue("password")
	cfmpwd := r.FormValue("confirmpassword")
	fmt.Println(Name)

	u := new(User)
	u.Username = Name
	u.Password = pwd
	u.confirmpassword = cfmpwd
	fmt.Println(*u)
	fmt.Println("method", r.Method)
	err = t.Execute(w, &u)
	if err != nil {
		log.Println("execution error", err)
	}
}

















/*
	var user models.User
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &user)

	mgosession :=DB.Obj.Copy()
	defer mgosession.Close()
	collection, err :=DB.DBconnect()
	if err != nil {
		log.Fatal(err)
	}


	var res models.ResponseResult
	err = collection.Find(bson.D{{"username", user.Username}})
	if err != nil {
		log.Println("err:", err)
	}
	doc := User{
		Username: r.Form["username"][0],
		Password: r.Form["password"][0],
	}
	err = collection.Insert(doc)
	if err != nil {
		panic(err)
	}

	if err != nil {
		res.Error = "Error While Creating User, Try Again"
		return
	}
	res.Result = "Registration Successful"
	return*/
		/*if r.Method == "GET" {
		fmt.Println("method", r.Method)
		t, err := template.ParseFiles("view/accountcontroller/createAccount.html")
		if err != nil {
			log.Println("error", err)
		}
		t.Execute(w, nil)
		r.ParseForm()
		n := r.Form.Get("username")
		fmt.Println(r)
		log.Println("CreateAccount Page")
		fmt.Println("username:", n)
		fmt.Println("password:", r.Form.Get("password"))
	} else {
		fmt.Println("method", r.Method)
		DB.DBconnect()
		//mgosession :=DB.Obj.Copy()
		//defer mgosession.Close()
		mgosession.SetMode(mgo.Monotonic, true)
		c := mgosession.DB("Webpage").C("login")
		doc := Signup{
			Username: r.Form["username"][0],
			Password: r.Form["password"][0],
		}
		err := c.Insert(doc)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%T\n", r.Form["username"])
		fmt.Println("---------------------------------------------")
		fmt.Println("username:", r.Form["username"][0]) //output:- username: user_name
		fmt.Println("password:", r.Form["password"][0]) //password:- password: user_password
	}

	r.ParseForm()
	uName, pwd, pwdConfirm := "", "", ""
	uName = r.FormValue("username")             // Data from the form
	pwd = r.FormValue("password")               // Data from the form
	pwdConfirm = r.FormValue("confirmpassword") // Data from the form

	// Empty data checking
	uNameCheck := models.IsEmpty(uName)
	pwdCheck := models.IsEmpty(pwd)
	pwdConfirmCheck := models.IsEmpty(pwdConfirm)

	if uNameCheck || pwdCheck || pwdConfirmCheck {
		fmt.Fprintf(w, "ErrorCode is -10 : There is empty data.")
		return
	}

	if pwd == pwdConfirm {
		// Save to database (username, email and password)
		fmt.Fprintln(w, "Registration successful.")
	} else {
		fmt.Fprintln(w, "Password information must be the same.")
	}
	_uName, _email, _pwd, _confirmPwd := false, false, false, false
	_uName = !models.IsEmpty(uName)
	_pwd = !models.IsEmpty(pwd)
	_confirmPwd = !models.IsEmpty(pwdConfirm)

	if _uName && _email && _pwd && _confirmPwd {
		fmt.Fprintln(w, "Username for Register : ", uName)
		fmt.Fprintln(w, "Password for Register : ", pwd)
		fmt.Fprintln(w, "ConfirmPassword for Register : ", pwdConfirm)
	} else {
		fmt.Fprintln(w, "This fields can not be blank!")
	}
}
/*	t,err:=template.ParseFiles("view/accountcontroller/createAccount.html")
	if err !=nil{
		log.Println("error",err)
	}
	t.Execute(w,nil)

var username = r.FormValue("username")
		var password = r.FormValue("password")
	session:=DB.Obj.Copy()
	defer session.Close()

	type Signup struct {
		Username    string `json:"username"`
		Password    string `json:"password"`
	} 
	res :=make(map[string]interface{})
	reg :=Signup{}
	err = session.DB("Webpage").C("login").Insert(reg)
	if err !=nil{
		log.Println("err",err)
	}

	log.Println("",username)

	log.Println("",password)
	res["name"]=Signup{Username: username,Password: password}
	res["message"]="MESSAGE"

	}*/

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