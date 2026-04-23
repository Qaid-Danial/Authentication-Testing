package main

import (
	"fmt"

	"github.com/rivo/tview"
)

func CreateWindow() *tview.Application {

	app := tview.NewApplication()

	var (
		username string
		password string
	)

	form := tview.NewForm()
	form.AddInputField("Username: ", "", 20, nil, func(text string) {
		username = text
	})
	form.AddInputField("Password", "pass", 20, nil,func(text string) {
		password = text
	})
	
	output := tview.NewTextView()
	output.SetBorder(true)

	form.AddButton("Input Validation", func() {
		output.SetText(fmt.Sprintf("Username: %s | Password: %s", username, password))
	})

	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow).AddItem(nil, 0, 1, false).AddItem(
		tview.NewFlex().AddItem(nil, 0, 1, false).AddItem(form, 50, 1, true).AddItem(nil, 0, 1, false), 10, 1, true).AddItem(output, 0, 1, false)


	app.SetRoot(flex, true)

	return app
}