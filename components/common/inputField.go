package common

import (
	"postgirl/color"

	"github.com/rivo/tview"
)

type InputFieldConfig struct {
	Placeholder string
	Label       string
	DefaultText string
	ChangedFunc func(text string)
}

func CreateInputField(cfg *InputFieldConfig) *tview.InputField {
	inputField := tview.NewInputField()
	inputField.SetPlaceholder(cfg.Placeholder)
	if cfg.ChangedFunc != nil {
		inputField.SetChangedFunc(func(text string) {
			if inputField.HasFocus() {
				cfg.ChangedFunc(text)
			}
		})
	}
	inputField.SetText(cfg.DefaultText)
	inputField.SetTitleColor(color.LABEL)

	if cfg.Label != "" {
		inputField.SetLabel(cfg.Label + " ")
	}

	inputField.SetPlaceholderTextColor(color.PLACEHOLDER)
	inputField.SetBackgroundColor(color.BACKGROUND_COMPONENT)
	inputField.SetFieldTextColor(color.LABEL)

	return inputField
}
