package components

import (
	"postgirl/color"
	"postgirl/components/common"
	"postgirl/lib"
	"postgirl/model"

	"github.com/rivo/tview"
)

func (c *Components) NewLayout() {
	var requestResponsePanel tview.Primitive

	if c.RequestResponsePanel != nil {
		requestResponsePanel = c.RequestResponsePanel.Root()
	} else {
		requestResponsePanel = common.CreateEmptyBox()
	}

	layoutFlex := tview.NewFlex()
	layoutFlex.AddItem(c.Sidebar.Root(), 30, 1, false)
	layoutFlex.AddItem(requestResponsePanel, 0, 1, false)
	layoutFlex.SetBackgroundColor(color.BACKGROUND)

	c.Layout = layoutFlex
}

func (c *Components) ChangePanel(req *model.Request) {
	if c.Layout.GetItemCount() == 2 {
		c.Layout.RemoveItem(c.Layout.GetItem(c.Layout.GetItemCount() - 1))
	}

	c.NewRequestResponsePanel(req)

	lib.Tview.UpdateDraw(func() {
		c.Layout.AddItem(c.RequestResponsePanel.Root(), 0, 1, false)
	})
}
