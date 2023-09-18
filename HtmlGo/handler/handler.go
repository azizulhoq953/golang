package handler

import (
	"fmt"
	"net/http"
	"text/template"
	"github/azizulhoq953/htmltemp/render"
)



func Home(w http.ResponseWriter, r *http.Request){ 
	/* w http.ResponseWriter construct and send an 
	HTTP response to the client.
	 You can use it to write the response body,
	  set HTTP headers, and manage the response */

	  render.RenderTemplate(w, "index-main.gohtml")

}

func About(w http.ResponseWriter, r *http.Request){
	
	 /*
	http.Request type contains information about the request, 
	such as the HTTP method (GET, POST, etc.), 
	request headers, query parameters, and the request body
	 */
	 render.RenderTemplate(w, "about.gohtml")
	           
}


type Contacts struct {
	UserName string
	Email string
	Message string
}


func Contact(w http.ResponseWriter, r *http.Request){


	tmpl :=template.Must(template.ParseFiles("templates/contact-2.gohtml"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	details2 :=  Contacts{
		UserName: r.FormValue("username"),
		Email: r.FormValue("email"),
		Message: r.FormValue("message"),
	}

	_=details2

	tmpl.Execute(w, struct {SuccessMessage bool} {true} )
	fmt.Println("User Messages:", details2)


	// render.RenderTemplate(w, "contact-2.gohtml")
}

type LoginData struct {
	Email string
	Password string
}



func Login(w http.ResponseWriter, r *http.Request){
	tmpl :=template.Must(template.ParseFiles("templates/login.gohtml"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	details :=  LoginData{
		Email: r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	_=details

	tmpl.Execute(w, struct {Success bool} {true} )
	fmt.Println("Login Informations", details)
}



	//render.RenderTemplate(w, "login.gohtml")
