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
		fmt.Println("Error when trying to opent he database: ", err)
		return nil 
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Error when trying to ping the server: ", err)
		return nil
	}

	fmt.Println("Database Connected Successfully")
	return db
}

func RegisterUser(db *sql.DB, username string, password string) {
	
	_, err := db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, password)

	if err != nil {
		fmt.Println("Error registering user: ", err)
		return
	}

	fmt.Println("New user was created!")
}

func ValidateUser(db *sql.DB, username string, password string) {

	var checkPass string

	if err := db.QueryRow("SELECT password FROM users WHERE username=$1", username).Scan(&checkPass); err != nil {
		fmt.Println("Error verifying user: ", err)
		return
	}

	if checkPass == password {
		fmt.Println("Password matched!")
	} else {
		fmt.Println("Password doesn't match!")
	}
}