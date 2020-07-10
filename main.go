package main

import (
	"awesomeProject1/controllers/accountcontroller"
	"log"
	"net/http"
)

/*type Page struct {
	Title string
	Body string
}

func PageHandler(w http.ResponseWriter, r *http.Request)  {
	a:= Page{Title: "Thoufeek", Body: "this my first page ever"}
	t,err:=template.ParseFiles("hometemplatepage.html")
	fmt.Println(err)
	t.Execute(w,a)

}
func handlerFunc(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w,"<h1>This is my page Thoufeek attempt1</h1>")

}*/
func main()  {

	http.HandleFunc("/account",accountcontroller.Index)
	http.HandleFunc("/account/index/login",accountcontroller.Login)
	http.HandleFunc("/account/index/login/welcome",accountcontroller.Welcome)
	err := http.ListenAndServe("localhost:8080", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: Not Listening..", err)
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
