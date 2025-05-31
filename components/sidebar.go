package components

import (
	"postgirl/color"
	"postgirl/components/common"
	"postgirl/internal/cache"
	"postgirl/internal/log"
	"postgirl/lib"
	"postgirl/model"
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
		if len(name) == 0 {
			common.ShowNotification(&common.NotificationConfig{
				Message: model.ErrNameRequired,
			})
			return
		}

		err := cache.CacheRequests.Create(name)
		if err != nil {
			common.ShowNotification(&common.NotificationConfig{
				Message: err.Error(),
			})
			return
		}

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

func (s *Sidebar) showModalLog() {
	content := common.CreateTextView(&common.TextViewConfig{
		Text: log.GetStringLogs(),
	})

	common.ShowModal(&common.ModalConfig{
		Content:     content,
		CloseButton: true,
		Center:      true,
		Width:       60,
		Height:      15,
		Title:       " 📜 Log ",
		TitleAlign:  tview.AlignCenter,
		BorderColor: color.BORDER,
	})
}

func (s *Sidebar) showModalInfo() {

	text := `Note: You can use arrow button for scroll this text info

[blue]URL[white]

[yellow]example[white] = http://example.com?foo=bar

[blue]PARAMS[white]

[yellow]example[white] = {
			"name": ["john doe"[]
		  }

[blue]HEADERS[white]
	
[yellow]example[white] = {
			"Custom-Header": "123"
		  }
	
[blue]BODY[white]

[yellow]type[white] = application/xml
[yellow]example[white] = <person>
			<name>
				john doe
			</name>
		  </person>

[yellow]type[white] = application/json, application/x-www-form-urlencoded, form-data
[yellow]example[white] = {
			"name": "john doe",
			[red]// specifically for send file on form-data[white]
			"image:file": ["avatar.jpg","product.png"]
		  }
	
[blue]FAQ[white]
[yellow]how to copy text from text area?[white]
You can select/block the text, and use combination ctrl + b
	`

	content := common.CreateTextView(&common.TextViewConfig{
		Border: false,
		Align:  tview.AlignLeft,
		Text:   text,
	})

	common.ShowModal(&common.ModalConfig{
		Content:     content,
		Title:       " ❗Info ",
		Width:       60,
		Height:      15,
		BorderColor: color.BORDER,
		CloseButton: true,
		TitleAlign:  tview.AlignCenter,
		Center:      true,
	})
}

func (s *Sidebar) showModalRemoveRequest() {
	selectedRequests := make(map[string]int)

	cacheLists := cache.CacheRequests.GetList()

	listRequest := common.CreateList(
		&common.ListConfig{
			Border: false,
			SetSelectedFunc: func(i int, s string, isCheckbox, checked bool) {
				if isCheckbox {
					if checked {
						selectedRequests[s] = i
					} else {
						delete(selectedRequests, s)
					}
				}
			},
		},
	)

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
		listRequest.AddItem(v, true)
	}

	flexContent := common.CreateFlex(&common.FlexConfig{
		Border:    false,
		Direction: tview.FlexRow,
	})
	flexContent.AddItem(listRequest.GetRoot(), 0, 1, false)
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
	s.list = list.GetRoot()

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
	},
	), 3, 1, false)
	actionsFlex.AddItem(common.CreateEmptyBox(), 1, 1, false)
	actionsFlex.AddItem(common.CreateButton(&common.ButtonConfig{
		Label:        "#",
		SelectedFunc: s.showModalLog,
	}), 3, 1, false)
	actionsFlex.AddItem(common.CreateEmptyBox(), 1, 1, false)

	sidebarFlex := common.CreateFlex(&common.FlexConfig{
		Border: false,
	})
	sidebarFlex.SetDirection(tview.FlexRow)
	sidebarFlex.AddItem(actionsFlex, 3, 1, false)
	sidebarFlex.AddItem(list.GetRoot(), 0, 1, false)

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
