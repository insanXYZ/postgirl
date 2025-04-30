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
	"strings"

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
	var selectedFile []string
	var list *common.List
	var modal *winman.WindowBase
	var fieldName string
	currentPath := "."

	ReadAndSetEntries := func(dirName string) error {
		list.RemoveAll()

		entries, err := util.ReadDir(dirName)
		if err != nil {
			return errors.New(model.ErrReadDirectory)
		}

		list.AddItem("↩ ..", false)

		for _, v := range entries {
			var icon string
			isDir := v.IsDir()

			if isDir {
				icon = "📁"
			} else {
				icon = "📄"
			}

			list.AddItem(icon+v.Name(), !isDir)
		}

		currentEntries = entries
		currentPath = dirName

		return nil
	}

	list = common.CreateList(&common.ListConfig{
		Border: false,
		SetSelectedFunc: func(i int, s string, isCheckbox, checked bool) {
			var name string
			var fullPath string

			if i == 0 {
				name = ".."
			} else {
				name = currentEntries[i-1].Name()
			}

			fullPath = filepath.Join(currentPath, name)

			if i == 0 || currentEntries[i-1].IsDir() {
				ReadAndSetEntries(fullPath)
			} else if currentEntries[i-1].Type().IsRegular() {
				if isCheckbox {
					if checked {
						selectedFile = append(selectedFile, fullPath)
					} else {
						ind := slices.Index(selectedFile, fullPath)
						selectedFile = slices.Delete(selectedFile, ind, ind+1)
					}
				}
			}
		},
	})

	fieldNameInput := common.CreateInputField(&common.InputFieldConfig{
		Label: "Fieldname",
		ChangedFunc: func(text string) {
			fieldName = text
		},
	})

	err := ReadAndSetEntries(currentPath)
	if err != nil {
		common.ShowNotification(&common.NotificationConfig{
			Message: err.Error(),
		})
	}

	addFileButton := common.CreateButton(&common.ButtonConfig{
		Label: "add file",
		SelectedFunc: func() {
			var m model.BodyMap

			trim := strings.TrimSpace(fieldName)

			if len(trim) == 0 {
				common.ShowNotification(&common.NotificationConfig{
					Message: model.ErrNameRequired,
				})
				return
			}

			err := util.JsonUnmarshal([]byte(a.BodyTextArea.GetText()), &m)
			if err != nil {
				common.ShowNotification(&common.NotificationConfig{
					Message: err.Error(),
				})
				return
			}

			m[fmt.Sprintf("%s:file", trim)] = selectedFile

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
		},
	})

	flex := common.CreateFlex(&common.FlexConfig{
		Direction: tview.FlexRow,
	})
	flex.AddItem(list.GetRoot(), 0, 1, true)
	flex.AddItem(fieldNameInput, 1, 1, false)
	flex.AddItem(common.CreateEmptyBox(), 1, 1, false)
	flex.AddItem(addFileButton, 1, 1, false)

	modal = common.ShowModal(&common.ModalConfig{
		Content:     flex,
		CloseButton: true,
		Center:      true,
		Width:       43,
		Height:      15,
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
	var dropdownBodyType *tview.DropDown

	attr := &Attribute{
		BodyTypeSelected: r.currentRequest.Attribute.BodyType,
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

	bodyTextArea := common.CreateTextArea(&common.TextAreaConfig{
		Border:       true,
		DefaultValue: r.currentRequest.Attribute.BodyString,
	})

	addFileButton := common.CreateButton(&common.ButtonConfig{
		Label: "Add file",
		SelectedFunc: func() {
			if !util.IsValidJson(bodyTextArea.GetText()) {
				common.ShowNotification(&common.NotificationConfig{
					Message: model.ErrInvalidFormatBody,
				})
				return
			}
			attr.ShowModalAddFile()
		},
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

	flexAttribute = common.CreateFlex(&common.FlexConfig{
		Border: false,
	})
	attr.Root = flexAttribute
	flexAttribute.SetDirection(tview.FlexRow)
	flexAttribute.AddItem(flexButton, 3, 1, false)
	flexAttribute.AddItem(paramsTextArea, 13, 1, false)

	removeBodyDropdown := func() {
		if flexButton.GetItemCount() > 6 {
			for i := flexButton.GetItemCount() - 1; i >= 6; i-- {
				flexButton.RemoveItem(flexButton.GetItem(i))
			}
		}
	}

	dropdownBodyType = common.CreateDropdown(&common.DropdownConfig{
		Options:        model.BodyOptions,
		CurrentOptions: slices.Index(model.BodyOptions, r.currentRequest.Attribute.BodyType),
	})
	dropdownBodyType.SetSelectedFunc(func(text string, index int) {
		attr.BodyTypeSelected = text

		removeBodyDropdown()

		if text == model.FORM_DATA {
			flexButton.AddItem(dropdownBodyType, 20, 1, false)
			flexButton.AddItem(common.CreateEmptyBox(), 1, 1, false)
			flexButton.AddItem(addFileButton, 15, 1, false)
		} else {
			flexButton.AddItem(dropdownBodyType, 20, 1, false)
		}
	})

	bodyButton.SetSelectedFunc(func() {
		lib.Tview.UpdateDraw(func() {
			if flexButton.GetItemCount() == 6 {
				flexButton.AddItem(dropdownBodyType, 20, 1, false)

				if attr.BodyTypeSelected == model.FORM_DATA {
					flexButton.AddItem(common.CreateEmptyBox(), 1, 1, false)
					flexButton.AddItem(addFileButton, 15, 1, false)
				}
			}

			if flexAttribute.GetItemCount() == 2 {
				flexAttribute.RemoveItem(flexAttribute.GetItem(flexAttribute.GetItemCount() - 1))
				flexAttribute.AddItem(bodyTextArea, 13, 1, false)
			}
		})
	})

	paramsButton.SetSelectedFunc(func() {
		removeBodyDropdown()
		lib.Tview.UpdateDraw(func() {
			flexAttribute.RemoveItem(flexAttribute.GetItem(flexAttribute.GetItemCount() - 1))
			flexAttribute.AddItem(paramsTextArea, 13, 1, false)
		})
	})

	headersButton.SetSelectedFunc(func() {
		removeBodyDropdown()
		lib.Tview.UpdateDraw(func() {
			flexAttribute.RemoveItem(flexAttribute.GetItem(flexAttribute.GetItemCount() - 1))
			flexAttribute.AddItem(headersTextArea, 13, 1, false)
		})
	})

	r.attribute = attr
}
