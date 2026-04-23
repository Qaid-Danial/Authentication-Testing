package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	TextMessage string `json:"textMessage"`
}

func receiveMsg(w http.ResponseWriter, r *http.Request) {
	var msg Message

	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

	fmt.Println(msg.TextMessage)

	reply := Message{
		TextMessage: "Hey Whats Good?",
	}

	marshalled, err := json.Marshal(reply)
	if err != nil {
		fmt.Println("Unable to marshal message: ", err)
		return 
	}

	w.Write(marshalled)
}

func main() {

	http.HandleFunc("/send", receiveMsg)

	fmt.Println("Server has started and is listening to port localhost:8080")
	http.ListenAndServe(":8080", nil)

}