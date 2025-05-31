package common

import (
	"postgirl/color"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"golang.design/x/clipboard"
)

type TextAreaConfig struct {
	DefaultValue string
	ChangedFunc  func(s string)
	Border       bool
}

func CreateTextArea(cfg *TextAreaConfig) *tview.TextArea {
	textArea := tview.NewTextArea()
	textArea.SetText(cfg.DefaultValue, false)
	textArea.SetBorder(cfg.Border)
	textArea.SetBorderColor(color.BORDER)
	textArea.SetBackgroundColor(color.BACKGROUND)
	textArea.SetTitleColor(tcell.ColorYellow)
	if cfg.ChangedFunc != nil {
		textArea.SetChangedFunc(func() {
			if textArea.HasFocus() {
				s := textArea.GetText()
				cfg.ChangedFunc(s)
			}
		})
	}
	textArea.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlB {
			s, _, _ := textArea.GetSelection()

			clipboard.Write(clipboard.FmtText, []byte(s))
		}
		return event
	})
	return textArea
}
