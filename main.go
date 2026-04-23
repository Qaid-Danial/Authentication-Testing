package main

import "fmt"

func main() {

	newWindow := CreateWindow()

	if err := newWindow.Run(); err != nil {
		fmt.Println("Unable to create new window: ", err)
		return
	}
}