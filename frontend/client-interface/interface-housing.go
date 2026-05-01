package clientinterface

import "github.com/rivo/tview"

func CreateWindow() (*tview.Application, string) {

	var tokenHousing string

	app := tview.NewApplication()

	pages := tview.NewPages()

	LoginPage, resetCall := LoginPage(func(username string, password string) {

		result, usertype, token := CheckCredentials(username, password)

		tokenHousing = token

		if result == "Success" {
			switch usertype{
			case "typeone":
				pages.SwitchToPage("Type1")
			case "typetwo":
				pages.SwitchToPage("Type2")
			}
		} else if result == "Unsuccess" {
		} else {
		}
	})

	logoutFunc := func() {
		pages.SwitchToPage("LoginPage")
		resetCall()
	}

	pages.AddPage("LoginPage", LoginPage, true, true)
	pages.AddPage("Type1", Type1Page(logoutFunc), true, false)
	pages.AddPage("Type2", Type2Page(logoutFunc), true, false)

	app.SetRoot(pages, true)
	app.EnableMouse(true)

	return app, tokenHousing
}