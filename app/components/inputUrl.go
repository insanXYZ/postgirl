package components

import (
	"net/http"

	"github.com/rivo/tview"
)

func (c *Components) NewInputUrl() {
	methodDropdown := tview.NewDropDown()
	methodDropdown.SetOptions([]string{
		http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodHead, http.MethodOptions,
	}, nil)

	input := tview.NewInputField()

	submitButton := tview.NewButton("send")

	flex := tview.NewFlex()
	flex.SetBorder(true)
	flex.AddItem(methodDropdown, 0, 1, false)
	flex.AddItem(input, 0, 1, false)
	flex.AddItem(submitButton, 5, 1, false)

	c.InputUrl = flex
}
