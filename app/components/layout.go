package components

import "github.com/rivo/tview"

func (c *Components) NewLayout() {
	editorPanel := tview.NewFlex()
	editorPanel.SetDirection(tview.FlexRow)
	editorPanel.AddItem(c.InputUrl, 3, 1, false)
	editorPanel.AddItem(c.Attribute, 13, 1, false)
	editorPanel.AddItem(c.Response, 0, 1, false)

	flex := tview.NewFlex()
	flex.AddItem(c.Sidebar, 30, 1, false)
	flex.AddItem(editorPanel, 0, 1, false)
	c.Layout = flex
}
