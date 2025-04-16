package components

import (
	"github.com/epiclabs-io/winman"
	"github.com/rivo/tview"
)

func (c *Components) NewSidebar() {

	showAddRequest := func() {
		content := tview.NewTextView().SetText("insan")

		wm := winman.NewWindow().Show()
		wm.SetTitle("✏️ add request")
		wm.SetRoot(content)
		wm.SetModal(true)
		wm.SetBorder(true)

		c.Winman.AddWindow(wm)
		c.Winman.Center(wm)

		go c.TviewApp.QueueUpdateDraw(func() {
			c.TviewApp.SetFocus(wm)
		})
	}

	list := tview.NewList()
	list.AddItem("item 1", "", 'a', nil)
	list.SetBorder(true)

	actionsFlex := tview.NewFlex()
	actionsFlex.AddItem(tview.NewTextView().SetText("postgirl"), 10, 1, false)
	actionsFlex.AddItem(tview.NewBox(), 0, 1, false)
	actionsFlex.AddItem(tview.NewButton("+").SetSelectedFunc(showAddRequest), 3, 1, false)
	actionsFlex.AddItem(tview.NewBox(), 1, 1, false)
	actionsFlex.AddItem(tview.NewButton("i"), 3, 1, false)
	actionsFlex.AddItem(tview.NewBox(), 1, 1, false)
	actionsFlex.SetBorder(true)

	sidebarFlex := tview.NewFlex()
	sidebarFlex.SetDirection(tview.FlexRow)
	sidebarFlex.AddItem(actionsFlex, 3, 1, false)
	sidebarFlex.AddItem(list, 0, 1, false)

	c.Sidebar = sidebarFlex
}
