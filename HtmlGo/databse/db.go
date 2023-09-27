
package database

import (
    "database/sql"
    // "fmt"
    // "net/http"
	// "github/azizulhoq953/htmltemp/render"
    _ "github.com/lib/pq"
    // "github.com/yourusername/project-folder/render"
)

type Database struct {
	*sql.DB
}

// User represents a user in the database.
type User struct {
    ID       int
    email string
    password string
}

// ... (Other code for initializing the database connection)

// FindUserByUsername queries the database to find a user by their username.
func InitializeDB() (*Database, error) {
    connStr := "user=postgres dbname=ht_db sslmode=disable" // Replace with your PostgreSQL connection string.
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }

    err = db.Ping()
    if err != nil {
        return nil, err
    }

    return &Database{db}, nil
}

// ... (Other database-related functions, e.g., RegisterUser, UpdateUser, etc.)
