package components

import "github.com/rivo/tview"

func (c *Components) NewLayout() {
	var editorPanel tview.Primitive

	if c.RequestResponsePanel != nil {
		editorPanel = c.RequestResponsePanel.Root()
	} else {
		editorPanel = tview.NewBox()
	}

	layoutFlex := tview.NewFlex()
	layoutFlex.AddItem(c.Sidebar.Root(), 30, 1, false)
	layoutFlex.AddItem(editorPanel, 0, 1, false)

	c.Layout = layoutFlex
}
