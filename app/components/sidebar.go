package components

import (
	"github.com/rivo/tview"
)

func (c *Components) NewSidebar() {
	list := tview.NewList()
	list.AddItem("item 1", "", 'a', nil)
	list.SetBorder(true)
	c.Sidebar = list
}
