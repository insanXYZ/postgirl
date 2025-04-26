package common

import (
	"postgirl/color"

	"github.com/rivo/tview"
)

type ButtonConfig struct {
	Label        string
	SelectedFunc func()
}

func CreateButton(cfg *ButtonConfig) *tview.Button {
	button := tview.NewButton(cfg.Label)
	button.SetBackgroundColor(color.BACKGROUND_COMPONENT)
	button.SetLabelColor(color.LABEL)
	button.SetTitleColor(color.LABEL)
	button.SetBackgroundColorActivated(color.BACKGROUND_COMPONENT_ACTIVE)
	button.SetSelectedFunc(cfg.SelectedFunc)
	return button
}
