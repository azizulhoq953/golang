package main

import(
	"database/sql"
	"fmt"
	_"github.com/lib/pq"

)

const (
	host ="localhost"
	port =5432
	user ="postgres"
	
	password="new_password"
	// dbname ="my_db"
	dbname = "registration_db"
)

func main()  {
	psqlconn := fmt.Sprintf("host =%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()
	
	// insertStmt := `insert into "user_login"("username", "password") values('azizulhoq', '4305')`
	// _, e:= db.Exec(insertStmt)
	// CheckError(e)

	// insertDynStmt:= `insert into "user_login"("username","password") values($1, $2)`
	// _,e = db.Exec(insertDynStmt, "azizulhoq","4305")

//new test

insertStmt := `insert into "sign_Up"("email", "phoneNimber","user_ID", "fullName","password") values('azizulhoq@gamil.com', '+880 1706257588', '123', 'Md Azizul Hoq','pass@4305')`
_, e:= db.Exec(insertStmt)
CheckError(e)

insertDynStmt:= `insert into "sign_Up"("email","phoneNimber","user_ID","fullName","password") values($1, $2,$3, $4, $5)`
_,e = db.Exec(insertDynStmt, "azizulhoq@gamil.com", "+880 1706257588", "123", "Md Azizul Hoq","pass@4305")


	// http.ListenAndServe(:8080)
	
}

func CheckError(err error){
	if err != nil{
		panic(err)
	}
	fmt.Println("database connected succesfull")
}
