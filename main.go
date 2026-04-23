package main

func main() {

	newWindow := CreateWindow()

	if err := newWindow.Run(); err != nil {
		panic(err)
	}

}