package common

import "github.com/rivo/tview"

type DropdownConfig struct {
	Options        []string
	SelectedFunc   func(text string, index int)
	CurrentOptions int
}

func CreateDropdown(config *DropdownConfig) *tview.DropDown {
	dropdown := tview.NewDropDown()
	dropdown.SetOptions(config.Options, config.SelectedFunc)
	dropdown.SetCurrentOption(config.CurrentOptions)
	return dropdown
}
