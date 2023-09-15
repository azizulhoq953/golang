package main

import(
	"fmt"
	"net/http"
	"github/azizulhoq953/htmltemp/handler"
)


const portNumber =":3000"
func main()  {

	fs :=http.FileServer(http.Dir("assets"))



	http.HandleFunc("/",handler.Home)
	http.HandleFunc("/about",handler.About)
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	fmt.Println(fmt.Sprintf("Application Running On %s", portNumber))
	 http.ListenAndServe(portNumber, nil)
	
}