package handler

import (
	"fmt"
	"net/http"
	"text/template"
	// "github/azizulhoq953/htmltemp/HtmlGo/databse"
	"github/azizulhoq953/htmltemp/HtmlGo/render"
	 
	// "github/azizulhoq953/htmltemp"
	
	// "main/contacts"
	// "log"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Contacts struct {

	UserName string
	Email string
	Message string
	
}

type Database struct {
	*render.Renderer
}


type LoginData struct {
	
	Email string
	Password string
	
}




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





func Login(databse *render.Renderer) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }

        email := r.FormValue("email")
        password := r.FormValue("Password")

        // Check the email and password against the database.
        // user, err := InitializeDB(db, email)
		user , err := Database.InitializeDB()
        if err !=  nil && user.Email == email {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return 
        }

        if user != nil && user.Password == password {
            // Authentication successful, redirect to a protected page.
            http.Redirect(w, r, "/about", http.StatusSeeOther)
            return
        }

        // Authentication failed, show an error message.
        fmt.Fprintf(w, "Authentication failed. Please try again.")
    }
}



	//render.RenderTemplate(w, "login.gohtml")
