package main

import(
	"fmt"
	"net/http"
	"text/template"
	
	// "github/azizulhoq953/htmltemp/form"
)

type Contacts struct {
	Email string
	Password string
	Subject string
	Messages string
}

const portNumber =":8080"
func main()  {

	// fs :=http.FileServer(http.Dir("assets"))
	tmpl :=template.Must(template.ParseFiles("forms.html"))


	// http.HandleFunc("/login",handler.Login)
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		details :=  Contacts{
			Email: r.FormValue("email"),
			Password: r.FormValue("password"),
			Subject: r.FormValue("subject"),
			Messages: r.FormValue("message"),
		}

		_=details

		tmpl.Execute(w, struct {Success bool} {true} )
	})

	// http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	fmt.Println(fmt.Sprintf("Application Running On %s", portNumber))
	http.ListenAndServe(portNumber, nil)


	 
	
}