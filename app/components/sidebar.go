package components

import "github.com/rivo/tview"

func NewSidebar() *tview.List {
	list := tview.NewList()
	list.AddItem("item 1", "", 'a', nil)
	return list
}
