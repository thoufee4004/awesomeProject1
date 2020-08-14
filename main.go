package main

import (
	"awesomeProject1/DB"
	"awesomeProject1/routers"
)


func main() {

	//db
	DB.DBconnect()
	//routers
	routers.Routeconfiguration()

	//Port
	//port()
	/*//path
	_, err := os.Stat(filepath.Join(".", "/assets/style.css", "style.css"))
	checkErr(err)*/
}



/*//Rverse
func Reverse{
	for i,j:=0,len(a)-1;i < j ;i,j=i+1,j-1{
		a[i],a[j]=a[j],a[i]
	}
}*/

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