package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB{
	connStr := "postgres://admin:pass@localhost:5432/authdb?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println("Unable to open the database: ", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Unable to ping the database: ", err)
		return nil
	}

	fmt.Println("Successfully connected to database!")
	return db
}

func Authenticator(db *sql.DB, username string, password string) string{

	var passwordCheck string

	if err := db.QueryRow("SELECT password FROM users WHERE username=$1", username).Scan(&passwordCheck); err != nil {
		return ("Error verifying user: " + string(err.Error()))
	}

	if password == passwordCheck {
		return "Sucess"
	} else {
		return "Password is wrong"
	}
	
}