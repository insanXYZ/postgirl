package common

import (
	"postgirl/color"

	"github.com/rivo/tview"
)

type FormConfig struct {
	Border bool
}

func CreateForm(cfg *FormConfig) *tview.Form {
	form := tview.NewForm()
	form.SetBackgroundColor(color.BACKGROUND)
	form.SetBorder(cfg.Border)
	form.SetBorderColor(color.BORDER)
	form.SetTitleColor(color.LABEL)
	form.SetButtonBackgroundColor(color.BACKGROUND_COMPONENT)
	form.SetButtonTextColor(color.LABEL)
	form.SetFieldBackgroundColor(color.BACKGROUND_COMPONENT)
	form.SetFieldTextColor(color.LABEL)
	return form
}
