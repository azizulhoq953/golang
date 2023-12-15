package main

import (
	"fmt"
	"net/http"

	// "text/template"
	"htmltemp/HtmlGo/handler"

	database "htmltemp/HtmlGo/databse"
	"log"
	// "github.com/yourusername/project-folder/render"
	// "github/azizulhoq953/htmltemp/form"
)

// type Contacts struct {
// 	Email string
// 	Password string
// }

const portNumber = ":8080"

func main() {

	fs := http.FileServer(http.Dir("assets"))
	// tmpl :=template.Must(template.ParseFiles("templates/login.gohtml"))

	// database, err := DB.InitDB()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer database.Close()
	database, err := database.InitializeDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	// Create a renderer.
	// red := handler.Contacts()

	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)
	http.HandleFunc("/contact", handler.Contact)
	// http.HandleFunc("/login", handler.Login)

	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	fmt.Println(fmt.Sprintf("Application Running On %s", portNumber))
	err = http.ListenAndServe(portNumber, nil)
	if err != nil {
		log.Fatal(err)
	}

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
	//start dbms

	// database, err := db.InitDB()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	defer database.Close()

	// 	// Create a renderer.
	// 	renderer := render.NewRenderer()

	// Register routes and start the server.
	// http.HandleFunc("/", render.Login(handler))
	// http.HandleFunc("/login", db.LoginSubmitHandler(database, renderer))
	// http.HandleFunc("/dashboard", render.DashboardHandler(renderer))

	// port := ":8080"
	// log.Printf("Server is running on %s...\n", port)
	// err = http.ListenAndServe(port, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
