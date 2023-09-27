package main

import (
    "database/sql"
	"net/http"
	// "GoWeb-PostgreSQL/dbms"

	"github.com/gorilla/mux"
	// Import the `pq` package with a preceding underscore since it is imported as a side
	// effect. The `pq` package is a GO Postgres driver for the `database/sql` package.
	_ "github.com/lib/pq"
)

func newRouter() *mux.Router {
	// Declare a router
	r := mux.NewRouter()
	// Declare static file directory
	staticFileDirectory := http.Dir("./static/")
	// Create static file server for our static files, i.e., .html, .css, etc
	staticFileServer := http.FileServer(staticFileDirectory)
	// Create file handler. Although the static files are placed inside `./static/` folder
	// in our local directory, it is served at the root (i.e., http://localhost:8080/)
	// when browsed in a browser. Hence, we need `http.StripPrefix` function to change the
	// file serve path.
	staticFileHandler := http.StripPrefix("/", staticFileServer)
	// Add static file handler to our router
	r.Handle("/", staticFileHandler).Methods("GET")
	// Add handler for `get` and `post` people functions
	r.HandleFunc("/person", getPersonHandler).Methods("GET")
	r.HandleFunc("/person", createPersonHandler).Methods("POST")

	return r
}

func main() {
	// Setup connection to our postgresql database
	connString := `user=postgres 
				   password=new_password
				   host=localhost
				   port=5432
				   dbname=user_database
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

	// Place our opened database into a `dbstruct` and assign it to `store` variable.
	// The `store` variable implements a `Store` interface. The `store` variable was
	// declared globally in `store.go` file.
	storeInfo = &dbStore{db: db}

	// Create router
	r := newRouter()

	

	// Listen to the port. Go server's default port is 8080.
	http.ListenAndServe(":8000", r)
}



// func main() {
//     // Initialize the router
//     gin.SetMode(gin.ReleaseMode)
//     router := gin.Default()
    
    
//     // router := gin.Default()
//     // Connect to PostgreSQL database
//     var err error
//     db, err = gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=new_dbpr password=new_password sslmode=disable")
//     if err != nil {
//         panic(err)
//     }
//     defer db.Close()

//     // AutoMigrate the User model to create the users table
//     db.AutoMigrate(&User{})

//     // Serve static files from the "static" directory
//     router.Static("/static", "/static/")
//     // router.HandleFunc("/login",loginHandle )
//     // staticFileDirectory := http.Dir("./static/")
//     // Define the login route
//     router.GET("/login", func(c *gin.Context) {
//         c.HTML(http.StatusOK, "login.html", nil)
//     })

//     router.POST("/login", func(c *gin.Context) {
//         // Retrieve form data
//         username := c.PostForm("username")
//         password := c.PostForm("password")

//         // Check credentials in the database
//         var user User
//         db.Where("username = ?", username).First(&user)

//         if user.Password == password {
//             // Successful login
//             c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
//         } else {
//             // Failed login
//             c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
//         }
//     })
    
    

    // Start the server
//     router.Run(":8888")
// }