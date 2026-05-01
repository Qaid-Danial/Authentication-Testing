package main

import clientinterface "auth-app/client-interface"

type Token struct {
	token string
}

var tkn Token

func main() {

	newWindow, token := clientinterface.CreateWindow()

	tkn.token = token

	if err := newWindow.Run(); err != nil {
		panic(err)
	}

}