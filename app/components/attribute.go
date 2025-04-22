package components

import (
	"postgirl/app/components/common"
	"postgirl/app/lib"
	"postgirl/app/model"
	"postgirl/app/util"

	"github.com/rivo/tview"
)

type Attribute struct {
	ParamsTextArea   *tview.TextArea
	HeadersTextArea  *tview.TextArea
	BodyTextArea     *tview.TextArea
	BodyTypeSelected string
	Root             *tview.Flex
}

// get value on headers , params and body text area
func (a *Attribute) GetText() (string, string, string) {
	return a.HeadersTextArea.GetText(), a.ParamsTextArea.GetText(), a.BodyTextArea.GetText()
}

func (a *Attribute) SetTextHeaders(v string) {
	a.HeadersTextArea.SetText(v, true)
}

func (a *Attribute) SetTextParams(v string) {
	a.ParamsTextArea.SetText(v, true)
}

func (r *RequestResponsePanel) NewAttribute() {
	var flexAttribute *tview.Flex
	attr := &Attribute{
		BodyTypeSelected: model.BodyOptions[0],
	}

	paramsButton := tview.NewButton("Params")
	paramsTextArea := common.CreateTextArea(&common.TextAreaConfig{
		Border: true,
	})
	attr.ParamsTextArea = paramsTextArea
	if stringParams, err := util.JsonMarshalString(r.currentModel.Attribute.Params); err == nil {
		paramsTextArea.SetText(stringParams, false)
	}

	headersButton := tview.NewButton("Headers")
	headersTextArea := common.CreateTextArea(&common.TextAreaConfig{
		Border: true,
	})
	attr.HeadersTextArea = headersTextArea
	if stringHeaders, err := util.JsonMarshalString(r.currentModel.Attribute.Headers); err == nil {
		headersTextArea.SetText(stringHeaders, false)
	}

	bodyButton := tview.NewButton("Body")

	dropdownBodyType := common.CreateDropdown(&common.DropdownConfig{
		Options:        model.BodyOptions,
		CurrentOptions: 0,
		SelectedFunc: func(text string, _ int) {
			attr.BodyTypeSelected = text
		},
	})

	bodyTextArea := common.CreateTextArea(&common.TextAreaConfig{
		Border: true,
	})

	attr.BodyTextArea = bodyTextArea

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
			if flexButton.GetItemCount() == 6 {
				flexButton.AddItem(dropdownBodyType, 15, 1, false)
			}

			if flexAttribute.GetItemCount() == 2 {
				flexAttribute.RemoveItem(flexAttribute.GetItem(flexAttribute.GetItemCount() - 1))
				flexAttribute.AddItem(bodyTextArea, 13, 1, false)
			}
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
