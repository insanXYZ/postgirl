package components

import "github.com/rivo/tview"

func NewResponse() *tview.Flex {
	responseMenu := tview.NewFlex()
	responseMenu.AddItem(tview.NewButton("Body"), 0, 1, false)
	responseMenu.AddItem(tview.NewButton("Headers"), 0, 1, false)

	responseTextArea := tview.NewTextArea()
	responseTextArea.SetBorder(true)

	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexColumn)
	flex.AddItem(responseMenu, 0, 1, false)
	flex.AddItem(responseTextArea, 0, 1, false)

	return flex
}
