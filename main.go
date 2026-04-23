package main

func main() {

	db := ConnectDB()

	// RegisterUser(db, "Danial", "0510")

	ValidateUser(db, "Danial", "1005")

	db.Close()
}