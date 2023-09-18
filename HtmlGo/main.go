package main

import(
	"fmt"
	"net/http"
	// "text/template"
	"github/azizulhoq953/htmltemp/handler"
	// "github/azizulhoq953/htmltemp/form"
)

// type Contacts struct {
// 	Email string
// 	Password string
// }

const portNumber =":8080"
func main()  {

	fs :=http.FileServer(http.Dir("assets"))
	// tmpl :=template.Must(template.ParseFiles("templates/login.gohtml"))

	http.HandleFunc("/",handler.Home)
	http.HandleFunc("/about",handler.About)
	http.HandleFunc("/contact",handler.Contact)
	http.HandleFunc("/login",handler.Login)
	// http.HandleFunc("/login",func(w http.ResponseWriter, r *http.Request){
	// 	if r.Method != http.MethodPost {
	// 		tmpl.Execute(w, nil)
	// 		return
	// 	}
	// 	details :=  Contacts{
	// 		Email: r.FormValue("email"),
	// 		Password: r.FormValue("password"),
	// 	}

	// 	_=details

	// 	tmpl.Execute(w, struct {Success bool} {true} )
	// 	fmt.Println("Login Informations", details)
	// })

	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	fmt.Println(fmt.Sprintf("Application Running On %s", portNumber))
	http.ListenAndServe(portNumber, nil)


	 
	
}