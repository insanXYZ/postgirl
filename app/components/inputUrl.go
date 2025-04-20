package components

import (
	"postgirl/app/components/common"
	"postgirl/app/model"

	"github.com/rivo/tview"
)

func (r *RequestResponsePanel) NewInputUrl() {
	methodDropdown := common.CreateDropdown(&common.DropdownConfig{
		Options: model.Methods,
		SelectedFunc: func(_ string, index int) {
			r.currentModel.Method = index
		},
		CurrentOptions: r.currentModel.Method,
	})

	inputUrl := common.CreateInputField(&common.InputFieldConfig{
		Placeholder: "Enter URL",
		DefaultText: r.currentModel.Url,
		ChangedFunc: func(text string) {
			r.currentModel.Url = text
		},
	})

	submitButton := tview.NewButton("send")
	submitButton.SetSelectedFunc(func() {
		r.submit <- true
	})

	flex := tview.NewFlex()
	flex.SetBorder(true)
	flex.AddItem(methodDropdown, 9, 1, false)
	flex.AddItem(inputUrl, 0, 1, false)
	flex.AddItem(tview.NewBox(), 2, 1, false)
	flex.AddItem(submitButton, 6, 1, false)

	r.inputUrl = flex
}
