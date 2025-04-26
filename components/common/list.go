package common

import (
	"postgirl/color"

	"github.com/rivo/tview"
)

type ListConfig struct {
	Border bool
}

func CreateList(cfg *ListConfig) *tview.List {
	list := tview.NewList()
	list.SetBorder(cfg.Border)
	list.SetBorderColor(color.BORDER)
	list.SetBackgroundColor(color.BACKGROUND)
	return list
}
