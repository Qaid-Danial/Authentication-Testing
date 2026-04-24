package main

import clientinterface "auth-app/client-interface"

func main() {

	newWindow := clientinterface.CreateWindow()

	if err := newWindow.Run(); err != nil {
		panic(err)
	}

}