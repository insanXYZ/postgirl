package components

import (
	"net/http"

	"github.com/rivo/tview"
)

func (e *EditorPanel) NewInputUrl() {
	methodDropdown := tview.NewDropDown()
	methodDropdown.SetOptions([]string{
		http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodHead, http.MethodOptions,
	}, func(text string, index int) {
		e.method = text
	})
	methodDropdown.SetCurrentOption(0)

	input := tview.NewInputField()
	input.SetChangedFunc(func(text string) {
		e.url = text
	})

	submitButton := tview.NewButton("send")
	submitButton.SetSelectedFunc(func() {
		e.submit <- true
	})

	flex := tview.NewFlex()
	flex.SetBorder(true)
	flex.AddItem(methodDropdown, 9, 1, false)
	flex.AddItem(input, 0, 1, false)
	flex.AddItem(tview.NewBox(), 1, 1, false)
	flex.AddItem(submitButton, 6, 1, false)

	e.inputUrl = flex
}
