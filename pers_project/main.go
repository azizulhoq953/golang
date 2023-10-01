package main

import (
	"fmt"
	// "FROM-HANDLING/db"
	"html/template"
	"database/sql"
	"net/http"
	// "GoWeb-PostgreSQL/dbms"

	// "github.com/gorilla/mux"
	// Import the `pq` package with a preceding underscore since it is imported as a side
	// effect. The `pq` package is a GO Postgres driver for the `database/sql` package.
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

type User struct {
	// id int
	Username string
	Email string
	Message string
	// Age int
	}

func newRouter()  {

	connString := `user=postgres 
	password=new_password
	host=localhost
	port=5432
	dbname=lin_database
	sslmode=disable`
db, err := sql.Open("postgres", connString)
if err != nil {
panic(err)
}

// Check whether we can access the database by pinging it
err = db.Ping()

if err != nil {
panic(err)
}
//  return 
}

// Place our opened database into a `dbstruct` and assign it to `store` variable.
// The `store` variable implements a `Store` interface. The `store` variable was
// declared globally in `store.go` file.
// storeInfo = &dbStore{db: db}

// Create router


// defer the close till after the main function has finished
// executing
// defer db.Close()


	// mysql








func main() {

	http.HandleFunc("/", home)
	//http.HandleFunc("/request", request)
	http.HandleFunc("/about", abouts)
	http.HandleFunc("/contact", contacts)
	http.HandleFunc("/login", logins)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":8000", nil)

	newRouter()
}



func home(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("templates/index-main.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)
	fmt.Fprintf(w, `welcome`)
}

func abouts(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("templates/about.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	// ptmp, err = ptmp.ParseFiles("wpage/features.gohtml")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	ptmp.Execute(w, nil)
	fmt.Fprintf(w, `welcome`)
}

func contacts(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("templates/contact-2.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	// ptmp, err = ptmp.ParseFiles("wpage/docs.gohtml")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	ptmp.Execute(w, nil)
	//fmt.Fprintf(w, `welcome`)
}

func logins(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("templates/login.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	// ptmp, err = ptmp.ParseFiles("wpage/docs.gohtml")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	ptmp.Execute(w, nil)
	//fmt.Fprintf(w, `welcome`)
}

// mysql
func request(w http.ResponseWriter, r *http.Request) {

	//method-1

	Username := r.FormValue("username")
	Email := r.FormValue("email")
	Message := r.FormValue("message")
	fmt.Println(Username, Email)                               //response
	fmt.Fprintf(w, `recived %s %s %s`, Username, Email, Message) //response

	//method-2
	// r.ParseForm()

	// for key, val := range r.Form {

	// fmt.Println(key, val)

	qs := "INSERT INTO `request` ( `username`, `email`, `message`) VALUES (NULL, '%s', '%s', '%s');"
	sql := fmt.Sprintf(qs, Username, Email, Message)
	insert, err := db.Query(sql)
	//if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()

	fmt.Fprintf(w, `ok`)
	//test for jenkins

}