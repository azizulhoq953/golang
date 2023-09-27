package main

import (
	// The sql go library is needed to interact with the database
	"database/sql"
)

// Store will have two methods, to add a new person, and to get all existing people
type Store interface {
	CreatePerson(person *Person) error
	GetPerson() ([]*Person, error)
}

// `dbStore` struct implements the `Store` interface. Variable `db` takes the pointer
// to the SQL database connection object.
type dbStore struct {
	db *sql.DB
}

// Create a global `store` variable of type `Store` interface. It will be initialized
// in `func main()`.
var storeInfo Store

func (storeInfo *dbStore) CreatePerson(person *Person) error {
	// 'Person' is a struct which has "nama", "birthday", and "occupation" attributes.
	// Type SQL query to insert new person into our database.
	// Note: `peopleinfo` is the name of the table within our `peopleDatabase` postgresql database.
	_, err := storeInfo.db.Query(
		"INSERT INTO user_loginpr(username,password) VALUES ($1,$2)",
		person.Username, person.Password)
	return err
}

func (storeInfo *dbStore) GetPerson() ([]*Person, error) {
	// Query the database for all persons, and return the result to the `rows` object.
	// Note: `peopleinfo` is the name of the table within our `peopleDatabase`
	rows, err := storeInfo.db.Query("SELECT username, password,  FROM user_loginpr")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create an empty slice of pointers to `Person` struct. This slice will be returned
	// by this function to its caller.
	personList := []*Person{}
	for rows.Next() {
		// For each row returned from the database, create a pointer to a `Person` struct.
		person := &Person{}
		// Populate the `Name`, `Birthday`, and `Occupation` attributes of the person
		if err := rows.Scan(&person.Username, &person.Password); err != nil {
			return nil, err
		}
		// Finally, append the new person to the returned slice, and repeat for the next row
		personList = append(personList, person)
	}
	return personList, nil
}