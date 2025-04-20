package components

import (
	"postgirl/app/components/common"
	"postgirl/app/lib"
	"postgirl/app/model"
	"postgirl/app/util"

	"github.com/rivo/tview"
)

type Attribute struct {
	ParamsTextArea  *tview.TextArea
	HeadersTextArea *tview.TextArea
	Root            *tview.Flex
}

func (r *RequestResponsePanel) NewAttribute() {
	var flexAttribute *tview.Flex
	attr := &Attribute{}

	paramsButton := tview.NewButton("Params")
	paramsTextArea := tview.NewTextArea()
	attr.ParamsTextArea = paramsTextArea
	paramsTextArea.SetBorder(true)
	if stringParams, err := util.Marshal(r.currentModel.Attribute.Params); err == nil {
		paramsTextArea.SetText(stringParams, false)
	}

	headersButton := tview.NewButton("Headers")
	headersTextArea := tview.NewTextArea()
	attr.HeadersTextArea = headersTextArea
	headersTextArea.SetBorder(true)
	if stringHeaders, err := util.Marshal(r.currentModel.Attribute.Headers); err == nil {
		headersTextArea.SetText(stringHeaders, false)
	}

	bodyButton := tview.NewButton("Body")

	flexButton := tview.NewFlex()
	flexButton.AddItem(paramsButton, 10, 1, false)
	flexButton.AddItem(common.CreateEmptyBox(), 1, 1, false)
	flexButton.AddItem(headersButton, 10, 1, false)
	flexButton.AddItem(common.CreateEmptyBox(), 1, 1, false)
	flexButton.AddItem(bodyButton, 10, 1, false)
	flexButton.AddItem(common.CreateEmptyBox(), 1, 1, false)
	flexButton.SetBorder(true)

	flexAttribute = tview.NewFlex()
	attr.Root = flexAttribute
	flexAttribute.SetDirection(tview.FlexRow)
	flexAttribute.AddItem(flexButton, 3, 1, false)
	flexAttribute.AddItem(common.CreateEmptyBox(), 13, 1, false)

	removedBodyDropdown := func() {
		if flexButton.GetItemCount() == 7 {
			lib.Tview.UpdateDraw(func() {
				flexButton.RemoveItem(flexButton.GetItem(flexButton.GetItemCount() - 1))
			})
		}
	}

	bodyButton.SetSelectedFunc(func() {
		lib.Tview.UpdateDraw(func() {

			dropdownBody := common.CreateDropdown(&common.DropdownConfig{
				Options:        model.BodyOptions,
				CurrentOptions: 0,
			})

			flexButton.AddItem(dropdownBody, 15, 1, false)
		})
	})

	paramsButton.SetSelectedFunc(func() {
		removedBodyDropdown()
		lib.Tview.UpdateDraw(func() {
			flexAttribute.RemoveItem(flexAttribute.GetItem(flexAttribute.GetItemCount() - 1))
			flexAttribute.AddItem(paramsTextArea, 13, 1, false)
		})
	})

	headersButton.SetSelectedFunc(func() {
		removedBodyDropdown()
		lib.Tview.UpdateDraw(func() {
			flexAttribute.RemoveItem(flexAttribute.GetItem(flexAttribute.GetItemCount() - 1))
			flexAttribute.AddItem(headersTextArea, 13, 1, false)
		})
	})

	r.attribute = attr

}
