package common

import (
	"postgirl/color"

	"github.com/rivo/tview"
)

type TextViewConfig struct {
	Border bool
	Text   string
	Align  int
}

func CreateTextView(cfg *TextViewConfig) *tview.TextView {
	textView := tview.NewTextView()
	textView.SetBorder(cfg.Border)
	textView.SetText(cfg.Text)
	textView.SetTextAlign(cfg.Align)
	textView.SetScrollable(true)
	textView.SetTextColor(color.LABEL)
	textView.SetDynamicColors(true)
	textView.SetBorderColor(color.BORDER)
	textView.SetWrap(false)
	textView.SetBackgroundColor(color.BACKGROUND)
	return textView
}
