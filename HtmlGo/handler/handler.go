package handler

import (
	// "fmt"
	"net/http"
	"github/azizulhoq953/htmltemp/render"
)


func Home(w http.ResponseWriter, r *http.Request){ 
	/* w http.ResponseWriter construct and send an 
	HTTP response to the client.
	 You can use it to write the response body,
	  set HTTP headers, and manage the response */

	  render.RenderTemplate(w, "main.gohtml")

}

func About(w http.ResponseWriter, r *http.Request){
	
	 /*
	http.Request type contains information about the request, 
	such as the HTTP method (GET, POST, etc.), 
	request headers, query parameters, and the request body
	 */
	 render.RenderTemplate(w, "about.gohtml")
	           
}