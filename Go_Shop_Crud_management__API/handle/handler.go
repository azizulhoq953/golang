package handle

import (
	"fmt"
	"html/template"
	"net/http"
)

type Renderer struct {
	templates *template.Template
}

func RenderTemplate(w http.ResponseWriter, gohtml string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + gohtml)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
	}
}

// func ServeHTML(w http.ResponseWriter, r *http.Request) {

// 	ptmp, err := template.ParseFiles("/templates/tmp/base.gohtml")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	ptmp.Execute(w, nil)
// 	fmt.Fprintf(w, `welcome`)

// }

func ServeHTML(w http.ResponseWriter, r *http.Request) {
	// Parse the HTML file

	tmpl := template.Must(template.ParseFiles("templates/base.gohtml"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	// tmpl, err := template.ParseFiles("./templates/tmp/main.gohtml")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// println(tmpl)

	// data := struct {
	// 	Title string
	// }{
	// 	Title: "My Page",
	// }

	// // Execute the template and write it to the response
	// if err := tmpl.Execute(w, data); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
}

// func main() {

// 	ServeHTML(w http.ResponseWriter, r*http.Request)
// }
