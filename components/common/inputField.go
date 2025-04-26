package common

import (
	"postgirl/color"

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

	inputField.SetPlaceholderTextColor(color.PLACEHOLDER)
	inputField.SetBackgroundColor(color.BACKGROUND_COMPONENT)
	inputField.SetFieldTextColor(color.LABEL)

	return inputField
}
