package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	db = ConnectDB()
	jwtKey = []byte("Hella_Secure_Key_Broski")
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Result struct {
	Result string `json:"result"`
	Usertype  string `json:"usertype"`
	Token string `json:"token"`
}

type DBResult struct {
	result string
	username string 
	usertype string
}

func checkCreds(w http.ResponseWriter, r *http.Request) {

	var userCred Credentials

	if err := json.NewDecoder(r.Body).Decode(&userCred); err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

	var respBody Result

	dbRespond := Authenticator(db, userCred.Username, userCred.Password)
	
	if dbRespond.result == "Success" {

		
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": dbRespond.username,
			"role": dbRespond.usertype,
			"exp": time.Now().Add(time.Hour).Unix(),
		})

		tokenString, _ := token.SignedString(jwtKey)


		fmt.Printf("\nUser %s successfully logged in | Token: %s", dbRespond.username, tokenString)

		respBody = Result{
			// Result: resp,
			// Usertype: usertype,
			Result: dbRespond.result,
			Usertype: dbRespond.usertype,
			Token: tokenString,
		}
	} else if dbRespond.result == "Unsuccess" {

		fmt.Printf("\nAn attempt to login on user %s was made, but unsuccessful!", userCred.Username)

		respBody = Result{
			Result: dbRespond.result,
			Usertype: dbRespond.usertype,
			Token: "",
		}
	} else {

		fmt.Println(dbRespond.result)

		respBody = Result{
			Result: "Invalid",
			Usertype: dbRespond.usertype,
			Token: "",
		}
	}

	marshalled, err := json.Marshal(respBody)
	if err != nil {
		fmt.Println("Unable to marshal message: ", err)
		return
	}

	w.Write(marshalled)
}

func main() {

	http.HandleFunc("/check-credentials", checkCreds)

	fmt.Println("Server has started and is listening to port localhost:8080")
	http.ListenAndServe(":8080", nil)

	

}