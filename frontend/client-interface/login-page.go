package clientinterface

import (
	"github.com/rivo/tview"
)

func LoginPage(SendCredentials func(username string, password string)) (tview.Primitive, func()) {
	var (
		username string
		password string
	)

	form := tview.NewForm()

	// form.AddInputField("Username: ", "", 20, nil, func(text string) {
	// 	username = text
	// })
	// form.AddInputField("Password", "", 20, nil,func(text string) {
	// 	password = text
	// })

	usernameInput := tview.NewInputField().SetLabel("Username: ").SetFieldWidth(20)
	passwordInput := tview.NewInputField().SetLabel("Passowrd: ").SetFieldWidth(20)

	usernameInput.SetChangedFunc(func(text string) {
		username = text
	})
	passwordInput.SetChangedFunc(func(text string) {
		password = text
	})

	form.AddFormItem(usernameInput)
	form.AddFormItem(passwordInput)

	form.AddButton("Input Validation", func() {
		if password == "" && username == "" {
			return
		}
		SendCredentials(username, password)
	})

	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow)
	flex.AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().
			AddItem(nil, 0, 1, false).
			AddItem(form, 50, 1, true).
			AddItem(nil, 0, 1, false), 10, 1, true).
		AddItem(nil, 0, 1, false)

	return flex, func ()  {

		username = ""
		password = ""

		usernameInput.SetText("")
		passwordInput.SetText("")
	}
}