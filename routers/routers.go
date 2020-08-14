package routers

import (
	"awesomeProject1/controllers/accountcontroller"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var Router = mux.NewRouter()

func Routeconfiguration(){
	//CSS file
	FileServerObj := http.FileServer(http.Dir("./assets/"))
	Router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/",FileServerObj))
	//http.Handle("/assets/", http.StripPrefix("/assets/", FileServerObj))

	//Routes
	Router.HandleFunc("/0", accountcontroller.Welcome)
	Router.HandleFunc("/1", accountcontroller.CreateAccount)
	Router.HandleFunc("/2", accountcontroller.Index)
	//Router.HandleFunc("/3{id}", accountcontroller.DeleteAccount).Methods("DELETE")

	http.Handle("/",Router)

	fmt.Println("listening...")
	err:= http.ListenAndServe(":2000", Router) // setting listening port
	if err!=nil{
		log.Println("err",err)
	}
}
func port() *mux.Router {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	//Listening server
	err := http.ListenAndServe(""+port, Router) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: Not Listening..", err)
	}
	log.Println("ListenAndServe: Listening", port)


	return Router
}
