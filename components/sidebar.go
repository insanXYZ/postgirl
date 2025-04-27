package components

import (
	"postgirl/color"
	"postgirl/components/common"
	"postgirl/internal/cache"
	"postgirl/lib"
	"postgirl/util"

	"github.com/epiclabs-io/winman"
	"github.com/rivo/tview"
)

type Sidebar struct {
	list *tview.List
	root *tview.Flex
}

func NewSidebar() *Sidebar {
	sidebar := Sidebar{}
	sidebar.NewList()

	if m, err := util.ReadCache(); err == nil {
		for i, v := range m {
			cache.CacheRequests.Create(i)
			cache.CacheRequests.SetRequest(i, &v)
			sidebar.AddList(i)
		}
	}
	return &sidebar
}

func (s *Sidebar) showModalAddRequest() {
	var name string
	var modal *winman.WindowBase

	form := common.CreateForm(&common.FormConfig{
		Border: false,
	})

	form.AddInputField("name", "", 0, nil, func(text string) {
		name = text
	})

	form.AddButton("create", func() {
		cache.CacheRequests.Create(name)
		s.AddList(name)
		common.RemoveModal(modal)

		SaveCache()
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

func (s *Sidebar) showModalInfo() {
}

func (s *Sidebar) showModalRemoveRequest() {
	selectedRequests := make(map[string]int)

	cacheLists := cache.CacheRequests.GetList()

	listRequest := common.CreateList(
		&common.ListConfig{
			Border: false,
		},
	)
	listRequest.SetSelectedFunc(func(i int, s1, s2 string, r rune) {
		runeLabel := []rune(s1)
		label := string(runeLabel[4:])

		if _, ok := selectedRequests[label]; ok {
			runeLabel[1] = ' '
			delete(selectedRequests, label)
		} else {
			runeLabel[1] = 'X'
			selectedRequests[label] = i
		}

		listRequest.SetItemText(i, string(runeLabel), "")
	})

	deleteButton := common.CreateButton(&common.ButtonConfig{
		Label: "delete",
		SelectedFunc: func() {
			var indexLists []int

			for i, v := range selectedRequests {
				cache.CacheRequests.DeleteMap(i)
				indexLists = append(indexLists, v)
			}

			if len(indexLists) != 0 {
				util.SortDesc(indexLists)

				for _, v := range indexLists {
					cache.CacheRequests.DeleteList(v)
					listRequest.RemoveItem(v)
					s.list.RemoveItem(v)
				}
			}

			selectedRequests = make(map[string]int)

			SaveCache()
		},
	})

	for _, v := range cacheLists {
		listRequest.AddItem("( ) "+v, "", 0, nil)
	}

	flexContent := common.CreateFlex(&common.FlexConfig{
		Border:    false,
		Direction: tview.FlexRow,
	})
	flexContent.AddItem(listRequest, 0, 1, false)
	flexContent.AddItem(common.CreateEmptyBox(), 1, 1, false)
	flexContent.AddItem(deleteButton, 1, 1, false)

	common.ShowModal(&common.ModalConfig{
		Content:     flexContent,
		CloseButton: true,
		Center:      true,
		Width:       40,
		Height:      12,
		Title:       " ❌ Remove request ",
		TitleAlign:  tview.AlignCenter,
	})
}

func (s *Sidebar) NewList() {
	list := common.CreateList(&common.ListConfig{
		Border: true,
	})
	s.list = list

	actionsFlex := common.CreateFlex(&common.FlexConfig{
		Border:    true,
		Direction: tview.FlexColumn,
	})
	actionsFlex.AddItem(common.CreateTextView(&common.TextViewConfig{
		Text: "postgirl",
	}), 10, 1, false)
	actionsFlex.AddItem(common.CreateEmptyBox(), 0, 1, false)
	actionsFlex.AddItem(common.CreateButton(&common.ButtonConfig{
		Label:        "+",
		SelectedFunc: s.showModalAddRequest,
	}), 3, 1, false)
	actionsFlex.AddItem(common.CreateEmptyBox(), 1, 1, false)
	actionsFlex.AddItem(common.CreateButton(&common.ButtonConfig{
		Label:        "-",
		SelectedFunc: s.showModalRemoveRequest,
	}), 3, 1, false)
	actionsFlex.AddItem(common.CreateEmptyBox(), 1, 1, false)
	actionsFlex.AddItem(common.CreateButton(&common.ButtonConfig{
		Label:        "i",
		SelectedFunc: s.showModalInfo,
	}), 3, 1, false)
	actionsFlex.AddItem(common.CreateEmptyBox(), 1, 1, false)

	sidebarFlex := common.CreateFlex(&common.FlexConfig{
		Border: false,
	})
	sidebarFlex.SetDirection(tview.FlexRow)
	sidebarFlex.AddItem(actionsFlex, 3, 1, false)
	sidebarFlex.AddItem(list, 0, 1, false)

	s.root = sidebarFlex
}

func (s *Sidebar) AddList(label string) {
	lib.Tview.UpdateDraw(func() {
		s.list.AddItem(label, "", 0, func() {
			switchRequestResponsePanelChan <- label
		})
	})
}

func (s *Sidebar) Root() tview.Primitive {
	return s.root
}
