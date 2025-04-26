package common

import (
	"math"
	"postgirl/color"
	"postgirl/lib"
	"time"

	"github.com/rivo/tview"
)

type NotificationConfig struct {
	Message string
}

func ShowNotification(cfg *NotificationConfig) {
	content := tview.NewTextView().SetText(cfg.Message)

	height := math.Ceil(float64(len(cfg.Message)) / 38)

	modal := ShowModal(&ModalConfig{
		Content:     content,
		X:           9999,
		Y:           9999,
		Width:       40,
		Height:      2 + int(height),
		Title:       " 🔔 ",
		BorderColor: color.ERROR,
	})

	time.AfterFunc(2000*time.Millisecond, func() {
		lib.Tview.UpdateDraw(func() {
			lib.Winman.RemoveWindow(modal)
			lib.Tview.SetFocus(nil)
		})
	})
}
