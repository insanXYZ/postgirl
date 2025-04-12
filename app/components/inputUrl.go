package components

import (
	"net/http"

	"github.com/rivo/tview"
)

func NewInputUrl() *tview.Flex {
	methodDropdown := tview.NewDropDown()
	methodDropdown.SetOptions([]string{
		http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodHead, http.MethodOptions,
	}, nil)

	input := tview.NewInputField()

	submitButton := tview.NewButton("send")

	flex := tview.NewFlex()
	flex.AddItem(methodDropdown, 0, 1, false)
	flex.AddItem(input, 0, 1, false)
	flex.AddItem(submitButton, 0, 1, false)

	return flex
}
