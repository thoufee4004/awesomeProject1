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

	//Routes
	Router.HandleFunc("/1", accountcontroller.CreateAccount).Methods("GET")
	Router.HandleFunc("/2", accountcontroller.Index).Methods("GET")
	//Router.HandleFunc("/3{id}", accountcontroller.DeleteAccount).Methods("DELETE")

	http.Handle("/",Router)

	log.Println("1")
	err:= http.ListenAndServe(":2000", Router) // setting listening port
	log.Println("2")
	if err!=nil{
		log.Println("err",err)
	}else {
		fmt.Println("listening...")
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
