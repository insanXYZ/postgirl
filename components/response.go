package components

import (
	"postgirl/color"
	"postgirl/components/common"
	"postgirl/lib"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Response struct {
	bodyTextArea   *tview.TextArea
	headerTextArea *tview.TextArea
	Menu           *tview.Flex
	StatusCode     string
	Loading        chan bool
	Root           *tview.Flex
}

func (r *Response) Reset() {
	lib.Tview.UpdateDraw(func() {
		r.SetBodyText("")
		r.SetHeaderText("")
	})
}

func (r *Response) ListenChan() {
	for load := range r.Loading {
		lib.Tview.UpdateDraw(func() {
			if r.Menu.GetItemCount() == 5 {
				r.Menu.RemoveItem(r.Menu.GetItem(4))
			}

			if load {
				r.Menu.AddItem(common.CreateTextView(&common.TextViewConfig{
					Text:   "loading...",
					Border: false,
				}), 10, 1, false)
			} else {
				if r.StatusCode != "" {
					var bgColor tcell.Color

					statusCode, _ := strconv.Atoi(r.StatusCode)

					if statusCode < 400 {
						bgColor = color.SUCCESS
					} else {
						bgColor = color.ERROR
					}

					tag := common.CreateTag(&common.TagConfig{
						Label:           r.StatusCode,
						BackgroundColor: bgColor,
					})

					r.Menu.AddItem(tag, 6, 1, false)
				}
			}
		})
	}

}

func (r *Response) SetBodyText(text string) {
	r.bodyTextArea.SetText(text, false)
}

func (r *Response) SetHeaderText(text string) {
	r.headerTextArea.SetText(text, false)
}

func (r *RequestResponsePanel) NewResponse() {
	response := &Response{
		Loading: make(chan bool),
	}
	go response.ListenChan()

	bodyButton := common.CreateButton(&common.ButtonConfig{
		Label: "Body",
	})
	headersButton := common.CreateButton(&common.ButtonConfig{
		Label: "Headers",
	})

	responseMenu := common.CreateFlex(&common.FlexConfig{
		Border:    true,
		Direction: tview.FlexColumn,
	})
	responseMenu.AddItem(bodyButton, 10, 1, false)
	responseMenu.AddItem(tview.NewBox(), 1, 1, false)
	responseMenu.AddItem(headersButton, 10, 1, false)
	responseMenu.AddItem(common.CreateEmptyBox(), 0, 1, false)
	response.Menu = responseMenu

	bodyTextArea := common.CreateTextArea(&common.TextAreaConfig{
		Border: true,
	})
	response.bodyTextArea = bodyTextArea

	headerTextArea := common.CreateTextArea(&common.TextAreaConfig{
		Border: true,
	})
	response.headerTextArea = headerTextArea

	flex := common.CreateFlex(&common.FlexConfig{
		Border: false,
	})
	flex.SetDirection(tview.FlexRow)
	flex.AddItem(responseMenu, 3, 1, false)
	flex.AddItem(bodyTextArea, 0, 1, false)

	response.Root = flex

	removeSecondItemFlex := func() {
		if flex.GetItemCount() == 2 {
			flex.RemoveItem(flex.GetItem(1))
		}
	}

	bodyButton.SetSelectedFunc(func() {
		lib.Tview.UpdateDraw(func() {
			removeSecondItemFlex()
			flex.AddItem(bodyTextArea, 0, 1, false)
		})
	})

	headersButton.SetSelectedFunc(func() {
		lib.Tview.UpdateDraw(func() {
			removeSecondItemFlex()
			flex.AddItem(headerTextArea, 0, 1, false)
		})
	})

	r.response = response
}
