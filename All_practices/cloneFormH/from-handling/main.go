package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/*var db *sql.DB
var err error

func init() {

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err = sql.Open("mysql", "root:4305Az@tcp(127.0.0.1:3306)/hosting_db")

	// if there is an error opening the connection, handle it

	nysql
	if err != nil {
		//panic(err.Error())
	}*/
// defer the close till after the main function has finished
// executing
//defer db.Close()

/*
	mysql
	fmt.Println("db connection successful")

}*/

func main() {

	http.HandleFunc("/", home)
	//http.HandleFunc("/request", request)
	http.HandleFunc("/features", features)
	http.HandleFunc("/docs", docs)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)
	//fmt.Fprintf(w, `welcome`)
}

func features(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("wpage/features.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp.Execute(w, nil)
	//fmt.Fprintf(w, `welcome`)
}

func docs(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("wpage/docs.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)
	//fmt.Fprintf(w, `welcome`)
}

/* mysql
func request(w http.ResponseWriter, r *http.Request) {

	//method-1

	name := r.FormValue("name")
	company := r.FormValue("company")
	email := r.FormValue("email")
	fmt.Println(name, company)                               //response
	fmt.Fprintf(w, `recived %s %s %s`, name, company, email) //response

	//method-2
	//r.ParseForm()

	//for key, val := range r.Form {

	//fmt.Println(key, val)

	qs := "INSERT INTO `request` (`id`, `name`, `company`, `email`, `status`) VALUES (NULL, '%s', '%s', '%s', '1');"
	sql := fmt.Sprintf(qs, name, company, email)
	//insert, err := db.Query(sql)
	//if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()

	fmt.Fprintf(w, `ok`)
	//test for jenkins

}*/
