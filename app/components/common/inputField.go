package common

import (
	"postgirl/app/color"

	"github.com/rivo/tview"
)

type InputFieldConfig struct {
	Placeholder string
	DefaultText string
	ChangedFunc func(text string)
}

func CreateInputField(config *InputFieldConfig) *tview.InputField {
	inputField := tview.NewInputField()
	inputField.SetPlaceholder(config.Placeholder)
	inputField.SetChangedFunc(config.ChangedFunc)
	inputField.SetText(config.DefaultText)
	inputField.SetBackgroundColor(color.BACKGROUND_COMPONENT)
	inputField.SetFieldBackgroundColor(color.BACKGROUND_COMPONENT)
	inputField.SetFieldTextColor(color.LABEL)

	return inputField
}
