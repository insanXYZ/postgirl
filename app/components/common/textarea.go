package common

import (
	"postgirl/app/color"

	"github.com/rivo/tview"
)

type TextAreaConfig struct {
	DefaultValue string
	Border       bool
}

func CreateTextArea(cfg *TextAreaConfig) *tview.TextArea {
	textArea := tview.NewTextArea()
	textArea.SetText(cfg.DefaultValue, true)
	textArea.SetBorder(cfg.Border)
	textArea.SetBorderColor(color.BORDER)

	return textArea
}
