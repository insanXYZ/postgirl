package common

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TagConfig struct {
	Label           string
	BackgroundColor tcell.Color
}

func CreateTag(cfg *TagConfig) *tview.TextView {
	textView := CreateTextView(&TextViewConfig{
		Border: false,
		Text:   cfg.Label,
		Align:  tview.AlignCenter,
	})
	textView.SetBackgroundColor(cfg.BackgroundColor)
	return textView
}
