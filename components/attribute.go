package components

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"postgirl/components/common"
	"postgirl/lib"
	"postgirl/model"
	"postgirl/util"
	"slices"

	"github.com/epiclabs-io/winman"
	"github.com/rivo/tview"
)

type Attribute struct {
	ParamsTextArea   *tview.TextArea
	HeadersTextArea  *tview.TextArea
	BodyTextArea     *tview.TextArea
	BodyTypeSelected string
	Root             *tview.Flex
}

func (a *Attribute) ShowModalAddFile() {
	var currentEntries []os.DirEntry
	var list *tview.List
	var modal *winman.WindowBase
	currentPath := "."

	RemoveValuesList := func() {
		for i := list.GetItemCount(); i >= 0; i-- {
			list.RemoveItem(i)
		}
	}

	ReadAndSetEntries := func(dirName string) error {
		RemoveValuesList()

		entries, err := util.ReadDir(dirName)
		if err != nil {
			return errors.New(model.ErrReadDir + ", error detail :" + err.Error())
		}

		list.AddItem("↩ ..", "", 0, nil)

		for _, v := range entries {
			var icon string

			if v.IsDir() {
				icon = "📁"
			} else {
				icon = "📄"
			}

			list.AddItem(fmt.Sprintf("%s %s", icon, v.Name()), "", 0, nil)
		}

		currentEntries = entries
		currentPath = dirName

		return nil
	}

	list = common.CreateList(&common.ListConfig{
		Border: false,
	})

	err := ReadAndSetEntries(currentPath)
	if err != nil {
		common.ShowNotification(&common.NotificationConfig{
			Message: err.Error(),
		})
	}

	list.SetSelectedFunc(func(i int, s1, s2 string, r rune) {
		var name string
		var path string

		if i == 0 {
			name = ".."
		} else {
			name = currentEntries[i-1].Name()
		}

		path = filepath.Join(currentPath, name)

		if i == 0 || currentEntries[i-1].IsDir() {
			ReadAndSetEntries(path)
		} else if currentEntries[i-1].Type().IsRegular() {
			var m model.BodyMap

			err := util.JsonUnmarshal([]byte(a.BodyTextArea.GetText()), &m)
			if err != nil {
				common.ShowNotification(&common.NotificationConfig{
					Message: err.Error(),
				})
				return
			}

			m["field:file"] = path

			s, err := util.JsonMarshalString(m)
			if err != nil {
				common.ShowNotification(&common.NotificationConfig{
					Message: err.Error(),
				})
				return
			}

			lib.Tview.UpdateDraw(func() {
				a.BodyTextArea.SetText(s, false)
			})

			common.RemoveModal(modal)
		}
	})

	modal = common.ShowModal(&common.ModalConfig{
		Content:     list,
		CloseButton: true,
		Center:      true,
		Width:       43,
		Height:      11,
		Title:       " 🔗 Add File ",
		TitleAlign:  tview.AlignCenter,
	})
}

// get value on headers , params and body text area
func (a *Attribute) GetText() (string, string, string) {
	return a.HeadersTextArea.GetText(), a.ParamsTextArea.GetText(), a.BodyTextArea.GetText()
}

func (a *Attribute) SetTextHeaders(v string) {
	a.HeadersTextArea.SetText(v, false)
}

func (a *Attribute) SetTextParams(v string) {
	a.ParamsTextArea.SetText(v, false)
}

func (r *RequestResponsePanel) NewAttribute() {
	var flexAttribute *tview.Flex
	attr := &Attribute{
		BodyTypeSelected: model.BodyOptions[0],
	}

	paramsButton := common.CreateButton(&common.ButtonConfig{
		Label: "Params",
	})
	paramsTextArea := common.CreateTextArea(&common.TextAreaConfig{
		Border: true,
	})
	attr.ParamsTextArea = paramsTextArea
	if stringParams, err := util.JsonMarshalString(r.currentRequest.Attribute.Params); err == nil {
		paramsTextArea.SetText(stringParams, false)
	}

	headersButton := common.CreateButton(&common.ButtonConfig{
		Label: "Headers",
	})
	headersTextArea := common.CreateTextArea(&common.TextAreaConfig{
		Border: true,
	})
	attr.HeadersTextArea = headersTextArea
	if stringHeaders, err := util.JsonMarshalString(r.currentRequest.Attribute.Headers); err == nil {
		headersTextArea.SetText(stringHeaders, false)
	}

	bodyButton := common.CreateButton(&common.ButtonConfig{
		Label: "Body",
	})

	addFileButton := common.CreateButton(&common.ButtonConfig{
		Label: "Add file",
		SelectedFunc: func() {
			attr.ShowModalAddFile()
		},
	})

	dropdownBodyType := common.CreateDropdown(&common.DropdownConfig{
		Options:        model.BodyOptions,
		CurrentOptions: slices.Index(model.BodyOptions, r.currentRequest.Attribute.BodyType),
		SelectedFunc: func(text string, _ int) {
			attr.BodyTypeSelected = text

			// if text == model.FORM_DATA {

			// }
		},
	})

	bodyTextArea := common.CreateTextArea(&common.TextAreaConfig{
		Border:       true,
		DefaultValue: r.currentRequest.Attribute.BodyString,
	})

	attr.BodyTextArea = bodyTextArea

	flexButton := common.CreateFlex(&common.FlexConfig{
		Border:    true,
		Direction: tview.FlexColumn,
	})
	flexButton.AddItem(paramsButton, 10, 1, false)
	flexButton.AddItem(common.CreateEmptyBox(), 1, 1, false)
	flexButton.AddItem(headersButton, 10, 1, false)
	flexButton.AddItem(common.CreateEmptyBox(), 1, 1, false)
	flexButton.AddItem(bodyButton, 10, 1, false)
	flexButton.AddItem(common.CreateEmptyBox(), 0, 1, false)
	flexButton.AddItem(addFileButton, 0, 1, false)

	flexAttribute = common.CreateFlex(&common.FlexConfig{
		Border: false,
	})
	attr.Root = flexAttribute
	flexAttribute.SetDirection(tview.FlexRow)
	flexAttribute.AddItem(flexButton, 3, 1, false)
	flexAttribute.AddItem(paramsTextArea, 13, 1, false)

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
				flexButton.AddItem(dropdownBodyType, 20, 1, false)
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
