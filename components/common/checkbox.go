package common

import (
	"postgirl/color"

	"github.com/rivo/tview"
)

type CheckboxConfig struct {
	Label       string
	ChangedFunc func(checked bool)
}

func CreateCheckbox(cfg *CheckboxConfig) *tview.Checkbox {
	cb := tview.NewCheckbox()
	cb.SetLabel(cfg.Label)
	cb.SetChangedFunc(cfg.ChangedFunc)
	cb.SetFieldBackgroundColor(color.BACKGROUND_COMPONENT)

	return cb
}
