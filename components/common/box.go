package common

import (
	"postgirl/color"

	"github.com/rivo/tview"
)

func CreateEmptyBox() *tview.Box {
	box := tview.NewBox()
	box.SetBackgroundColor(color.BACKGROUND)
	return box
}
