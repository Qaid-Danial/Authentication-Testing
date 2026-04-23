package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)


type Message struct {
	TextMessage string `json:"textMessage"`
}


func ClientFunc() string {

	var msg Message
	var receive Message

	Reader := bufio.NewReader(os.Stdin)

	fmt.Print("> ")
	read, _ := Reader.ReadString('\n')

	msg = Message{
		TextMessage: read,
	}

	marshalled, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Unable to marshal: ", err)
	}

	resp, err := http.Post(
		"http://localhost:8080/send",
		"application/json",
		bytes.NewBuffer(marshalled),
	)
	if err != nil {
		fmt.Println("Unable to post: ", err)
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&receive); err != nil {
		fmt.Println("Unable to decode message: ", err)
	}

	return ("Message from server: " + receive.TextMessage)

}


type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SendCredentials(username string, password string) {

	cred := Credentials{
		Username: username,
		Password: password,
	}

	marshalled, err := json.Marshal(cred)
	if err != nil {
		fmt.Println("Unable to marshal: ", err)
	}

	resp, err := http.Post(
		"http://localhost:8080/check-credentials",
		"application/json",
		bytes.NewBuffer(marshalled),
	)
	if err != nil {
		fmt.Println("Unable to post: ", err)
	}

	defer resp.Body.Close()
}