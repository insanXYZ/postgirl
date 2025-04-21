package components

import (
	"postgirl/app/color"
	"postgirl/app/components/common"
	"postgirl/app/internal/cache"
	"postgirl/app/lib"

	"github.com/epiclabs-io/winman"
	"github.com/rivo/tview"
)

type Sidebar struct {
	list *tview.List
	root *tview.Flex

	//hook change panel
	requestResponsePanelChan chan string
}

func (c *Components) NewSidebar() {
	sidebar := Sidebar{
		requestResponsePanelChan: c.RequestResponsePanelChan,
	}
	sidebar.NewList()

	c.Sidebar = &sidebar
}

func (s *Sidebar) NewList() {

	showModalAddRequest := func() {
		var name string
		var modal *winman.WindowBase

		form := tview.NewForm()

		form.AddInputField("name", "", 0, nil, func(text string) {
			name = text
		})

		form.AddButton("create", func() {
			cache.CacheRequests.Create(name)
			s.AddList(name)
			common.RemoveModal(modal)
		})

		modal = common.ShowModal(&common.ModalConfig{
			Content:     form,
			Title:       " ➕ Add request ",
			Width:       40,
			Height:      7,
			BorderColor: color.BORDER,
			CloseButton: true,
			TitleAlign:  tview.AlignCenter,
			Center:      true,
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
		s.list.AddItem(label, "", '⦔', func() {
			s.requestResponsePanelChan <- label
		})
	})
}

func (s *Sidebar) Root() tview.Primitive {
	return s.root
}
