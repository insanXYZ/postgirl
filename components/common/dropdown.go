package common

import (
	"postgirl/color"

	"github.com/rivo/tview"
)

type DropdownConfig struct {
	Options        []string
	SelectedFunc   func(text string, index int)
	CurrentOptions int
	Width          int
}

func CreateDropdown(config *DropdownConfig) *tview.DropDown {
	dropdown := tview.NewDropDown()
	dropdown.SetFieldWidth(config.Width)
	dropdown.SetOptions(config.Options, config.SelectedFunc)
	dropdown.SetCurrentOption(config.CurrentOptions)
	dropdown.SetFieldTextColor(color.LABEL)
	dropdown.SetBackgroundColor(color.BACKGROUND_COMPONENT)
	return dropdown
}
