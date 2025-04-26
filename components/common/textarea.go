package common

import (
	"postgirl/color"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TextAreaConfig struct {
	DefaultValue string
	Border       bool
}

func CreateTextArea(cfg *TextAreaConfig) *tview.TextArea {
	textArea := tview.NewTextArea()
	textArea.SetText(cfg.DefaultValue, false)
	textArea.SetBorder(cfg.Border)
	textArea.SetBorderColor(color.BORDER)
	textArea.SetBackgroundColor(color.BACKGROUND)
	textArea.SetTitleColor(tcell.ColorYellow)

	return textArea
}
