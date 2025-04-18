package components

import (
	"fmt"
	"postgirl/app/lib"

	"github.com/rivo/tview"
)

func (r *RequestResponsePanel) NewAttribute() {
	var flexAttribute *tview.Flex

	paramsTextArea := tview.NewTextArea()
	var stringParams string
	if len(r.currentModel.Attribute.Headers) > 0 {
		for i, v := range r.currentModel.Attribute.Params {
			stringParams += fmt.Sprintf("%s:%s\n", i, v)
		}
	}
	paramsTextArea.SetText(stringParams, false)

	headersTextArea := tview.NewTextArea()
	var stringHeaders string
	if len(r.currentModel.Attribute.Headers) > 0 {
		for i, v := range r.currentModel.Attribute.Headers {
			stringHeaders += fmt.Sprintf("%s:%s\n", i, v)
		}
	}
	headersTextArea.SetText(stringHeaders, false)

	paramsButton := tview.NewButton("Params")
	headersButton := tview.NewButton("Headers")
	bodyButton := tview.NewButton("Body")

	flexButton := tview.NewFlex()
	flexButton.AddItem(paramsButton, 10, 1, false)
	flexButton.AddItem(tview.NewBox(), 1, 1, false)
	flexButton.AddItem(headersButton, 10, 1, false)
	flexButton.AddItem(tview.NewBox(), 1, 1, false)
	flexButton.AddItem(bodyButton, 10, 1, false)
	flexButton.AddItem(tview.NewBox(), 1, 1, false)
	flexButton.SetBorder(true)

	flexAttribute = tview.NewFlex()
	flexAttribute.SetDirection(tview.FlexRow)
	flexAttribute.AddItem(flexButton, 3, 1, false)
	flexAttribute.AddItem(tview.NewBox().SetBorder(true), 10, 1, false)

	removedBodyDropdown := func() {
		if flexButton.GetItemCount() == 7 {
			lib.Tview.UpdateDraw(func() {
				flexButton.RemoveItem(flexButton.GetItem(flexButton.GetItemCount() - 1))
			})
		}
	}

	bodyButton.SetSelectedFunc(func() {
		lib.Tview.UpdateDraw(func() {
			dropdownBody := tview.NewDropDown()

			dropdownBody.SetOptions(
				[]string{"none",
					"form-data",
					"x-www-form-urlencoded",
					"json",
					"xml"},

				func(text string, index int) {

				},
			)

			flexButton.AddItem(dropdownBody, 15, 1, false)
		})
	})

	paramsButton.SetSelectedFunc(func() {
		removedBodyDropdown()
		lib.Tview.UpdateDraw(func() {
			flexAttribute.RemoveItem(flexAttribute.GetItem(flexAttribute.GetItemCount() - 1))
			flexAttribute.AddItem(paramsTextArea, 10, 1, false)
		})
	})

	headersButton.SetSelectedFunc(func() {
		removedBodyDropdown()
		lib.Tview.UpdateDraw(func() {
			flexAttribute.RemoveItem(flexAttribute.GetItem(flexAttribute.GetItemCount() - 1))
			flexAttribute.AddItem(headersTextArea, 10, 1, false)
		})
	})

	r.attribute = flexAttribute
}
