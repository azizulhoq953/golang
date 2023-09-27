package render

import(
	"fmt"
	"html/template"
	"net/http"
)


type Renderer struct {
	templates *template.Template
}

// NewRenderer creates a new Renderer.
// func NewRenderer() *Renderer {
// 	templates := template.Must(template.ParseGlob("templates/*.html"))
// 	return &Renderer{templates}
// }

// LoginHandler handles rendering the login page.
// func (r *Renderer) LoginHandler() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// Render the login page.
// 		r.templates.ExecuteTemplate(w, "login.html", nil)
// 	}
// }

// // DashboardHandler handles rendering the dashboard page.
// func (r *Renderer) DashboardHandler() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// Render the dashboard page.
// 		r.templates.ExecuteTemplate(w, "dashboard.html", nil)
// 	}
// }



func RenderTemplate(w http.ResponseWriter, gohtml string){
	parsedTemplate, _:=template.ParseFiles("./templates/"+gohtml)
	err:=parsedTemplate.Execute(w, nil)
	if err != nil{
		fmt.Println("error parsing template", err)
	}
}