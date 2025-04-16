package components

import "github.com/rivo/tview"

func (c *Components) NewResponse() {

	bodyButton := tview.NewButton("Body")
	headersButton := tview.NewButton("Headers")

	responseMenu := tview.NewFlex()
	responseMenu.AddItem(bodyButton, 10, 1, false)
	responseMenu.AddItem(tview.NewBox(), 1, 1, false)
	responseMenu.AddItem(headersButton, 10, 1, false)
	responseMenu.SetBorder(true)

	responseTextArea := tview.NewTextArea()
	responseTextArea.SetBorder(true)
	responseTextArea.SetText("Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.", false)

	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow)
	flex.AddItem(responseMenu, 3, 1, false)
	flex.AddItem(responseTextArea, 0, 1, false)

	c.Response = flex
}
