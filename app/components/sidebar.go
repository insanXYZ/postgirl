package components

import (
	"fmt"
	"postgirl/app/components/util"
	"postgirl/app/lib"

	"github.com/rivo/tview"
)

type Sidebar struct {
	list *tview.List
	root *tview.Flex
}

func (c *Components) NewSidebar() {
	sidebar := Sidebar{}
	sidebar.NewList()

	c.Sidebar = &sidebar
}

func (s *Sidebar) NewList() {

	showModalAddRequest := func() {
		var name string

		form := tview.NewForm()
		form.AddInputField("name", "", 0, nil, func(text string) {
			name = text
		})
		form.AddButton("create", func() {
			lib.Tview.Stop()
			fmt.Println(name)
		})

		util.ShowModal(&util.ModalConfig{
			Content: form,
			Title:   "add request",
			Width:   40,
			Height:  7,
		})

	}

	list := tview.NewList()
	list.SetBorder(true)
	s.list = list

	actionsFlex := tview.NewFlex()
	actionsFlex.AddItem(tview.NewTextView().SetText("postgirl"), 10, 1, false)
	actionsFlex.AddItem(tview.NewBox(), 0, 1, false)
	actionsFlex.AddItem(tview.NewButton("+").SetSelectedFunc(showModalAddRequest), 3, 1, false)
	actionsFlex.AddItem(tview.NewBox(), 1, 1, false)
	actionsFlex.AddItem(tview.NewButton("i"), 3, 1, false)
	actionsFlex.AddItem(tview.NewBox(), 1, 1, false)
	actionsFlex.SetBorder(true)

	sidebarFlex := tview.NewFlex()
	sidebarFlex.SetDirection(tview.FlexRow)
	sidebarFlex.AddItem(actionsFlex, 3, 1, false)
	sidebarFlex.AddItem(list, 0, 1, false)

	s.root = sidebarFlex
}

func (s *Sidebar) AddList(label string) {
	lib.Tview.UpdateDraw(func() {
		s.list.AddItem(label, "", ' ', nil)
	})
}

func (s *Sidebar) Root() tview.Primitive {
	return s.root
}
