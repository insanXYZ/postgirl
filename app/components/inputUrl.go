package components

import (
	"postgirl/app/model"

	"github.com/rivo/tview"
)

func (r *RequestResponsePanel) NewInputUrl() {
	methodDropdown := tview.NewDropDown()
	methodDropdown.SetOptions(model.Methods, func(text string, index int) {
		r.currentModel.Method = index
	})
	methodDropdown.SetCurrentOption(r.currentModel.Method)

	inputUrl := tview.NewInputField()
	inputUrl.SetText(r.currentModel.Url)
	inputUrl.SetChangedFunc(func(text string) {
		r.currentModel.Url = text
	})

	submitButton := tview.NewButton("send")
	submitButton.SetSelectedFunc(func() {
		r.submit <- true
	})

	flex := tview.NewFlex()
	flex.SetBorder(true)
	flex.AddItem(methodDropdown, 9, 1, false)
	flex.AddItem(inputUrl, 0, 1, false)
	flex.AddItem(tview.NewBox(), 1, 1, false)
	flex.AddItem(submitButton, 6, 1, false)

	r.inputUrl = flex
}
