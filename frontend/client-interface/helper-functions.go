package clientinterface

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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

func CheckCredentials(username string, password string) (result string, usertype string, token string) {

	cred := Credentials{
		Username: username,
		Password: password,
	}

	marshalled, err := json.Marshal(cred)
	if err != nil {
		fmt.Println("Unable to marshal: ", err)
	}

	resp, err := http.Post(
		"https://go.silly-danial.lol/check-credentials",
		"application/json",
		bytes.NewBuffer(marshalled),
	)
	if err != nil {
		fmt.Println("Unable to post: ", err)
	}

	defer resp.Body.Close()

	var res Result

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		fmt.Println("Unable to decode message: ", err)
		return
	}

	return res.Result, res.Usertype, res.Token
}

