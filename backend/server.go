package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// type Message struct {
// 	TextMessage string `json:"textMessage"`
// }

// func receiveMsg(w http.ResponseWriter, r *http.Request) {
// 	var msg Message

// 	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
// 		http.Error(w, "Invalid Json", http.StatusBadRequest)
// 		return
// 	}

// 	fmt.Println(msg.TextMessage)

// 	reply := Message{
// 		TextMessage: "Hey Whats Good?",
// 	}

// 	marshalled, err := json.Marshal(reply)
// 	if err != nil {
// 		fmt.Println("Unable to marshal message: ", err)
// 		return
// 	}

// 	w.Write(marshalled)
// }

var (
	db = ConnectDB()
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Result struct {
	Result string `json:"result"`
	Usertype  string `json:"usertype"`
}

func checkCreds(w http.ResponseWriter, r *http.Request) {

	var userCred Credentials

	if err := json.NewDecoder(r.Body).Decode(&userCred); err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

	var respBody Result

	resp, usertype := Authenticator(db, userCred.Username, userCred.Password)
	
	if resp == "Success" {

		fmt.Printf("\nUser %s successfully logged in", userCred.Username)

		respBody = Result{
			Result: resp,
			Usertype: usertype,
		}
	} else if resp == "Unsuccess" {

		fmt.Printf("\nAn attempt to login on user %s was made, but unsuccessful!", userCred.Username)

		respBody = Result{
			Result: resp,
			Usertype: usertype,
		}
	} else {

		fmt.Println(resp)

		respBody = Result{
			Result: "Invalid",
			Usertype: usertype,
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

	// http.HandleFunc("/send", receiveMsg)
	http.HandleFunc("/check-credentials", checkCreds)

	fmt.Println("Server has started and is listening to port localhost:8080")
	http.ListenAndServe(":8080", nil)

}