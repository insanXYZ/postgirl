package common

import (
	"postgirl/color"

	"github.com/rivo/tview"
)

type FlexConfig struct {
	Border    bool
	Direction int
}

func CreateFlex(cfg *FlexConfig) *tview.Flex {
	flex := tview.NewFlex()
	flex.SetDirection(cfg.Direction)
	flex.SetBorder(cfg.Border)
	flex.SetBorderColor(color.BORDER)
	flex.SetBackgroundColor(color.BACKGROUND)
	return flex
}
