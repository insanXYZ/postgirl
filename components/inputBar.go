package components

import (
	"postgirl/components/common"
	"postgirl/model"

	"github.com/rivo/tview"
)

type InputBar struct {
	Method   int
	Dropdown *tview.DropDown

	Url      string
	InputUrl *tview.InputField

	Root *tview.Flex
}

func (r *RequestResponsePanel) NewInputUrl() {
	inputBar := &InputBar{}

	methodDropdown := common.CreateDropdown(&common.DropdownConfig{
		Options: model.Methods,
		Width:   11,
		SelectedFunc: func(_ string, index int) {
			inputBar.Method = index
		},
		CurrentOptions: r.currentRequest.Method,
	})
	inputBar.Dropdown = methodDropdown

	inputUrl := common.CreateInputField(&common.InputFieldConfig{
		Placeholder: "Enter URL",
		DefaultText: r.currentRequest.Url,
		ChangedFunc: func(text string) {
			inputBar.Url = text
		},
	})
	inputBar.InputUrl = inputUrl

	sendButton := common.CreateButton(&common.ButtonConfig{
		Label: "send",
		SelectedFunc: func() {
			r.send <- true
		},
	})

	flex := common.CreateFlex(&common.FlexConfig{
		Border:    true,
		Direction: tview.FlexColumn,
	})
	flex.AddItem(methodDropdown, 11, 1, false)
	flex.AddItem(common.CreateEmptyBox(), 2, 1, false)
	flex.AddItem(inputUrl, 0, 1, false)
	flex.AddItem(common.CreateEmptyBox(), 2, 1, false)
	flex.AddItem(sendButton, 6, 1, false)

	inputBar.Root = flex
	r.inputBar = inputBar
}

func (r *RequestResponsePanel) listenChan() {
	for range r.send {
		r.HandlerSend()
	}
}
