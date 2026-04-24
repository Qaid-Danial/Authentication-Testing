package clientinterface

import (
	"fmt"

	"github.com/rivo/tview"
)

func Type2Page(logout func()) *tview.Flex {

	box := tview.NewFlex()
	box.SetDirection(tview.FlexRow)

	panel := tview.NewTextView()
	panel.SetText(fmt.Sprintln("This is User Type 2 Page"))
	panel.SetBorder(true)

	button := tview.NewButton("Log Out").SetSelectedFunc(logout)
	button.SetBorder(true)

	box.AddItem(panel, 0, 1, false)
	box.AddItem(button, 3, 1, true)
	
	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow)

	flex.AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().
			AddItem(nil, 0, 1, false).
			AddItem(box, 50, 1, true).
			AddItem(nil, 0, 1, false), 10, 1, false).
		AddItem(nil, 0, 1, false)

	return flex
}