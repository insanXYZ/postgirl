package common

import (
	"postgirl/app/color"
	"postgirl/app/lib"

	"github.com/epiclabs-io/winman"
	"github.com/rivo/tview"
)

type ModalConfig struct {
	Content       tview.Primitive
	CloseFocus    tview.Primitive
	Width, Height int
	Title         string
}

func ShowModal(cfg *ModalConfig) *winman.WindowBase {
	wm := lib.Winman.NewWindow().Show()
	wm.SetTitle(cfg.Title)
	wm.SetRoot(cfg.Content)
	wm.SetModal(true)
	wm.SetBorder(true)
	wm.SetBackgroundColor(color.BACKGROUND)
	wm.SetTitleColor(color.BORDER)
	wm.SetBorderColor(color.BORDER)
	wm.SetRect(1, 1, cfg.Width, cfg.Height)
	wm.AddButton(&winman.Button{
		Symbol: 'X',
		OnClick: func() {
			lib.Winman.RemoveWindow(wm)
			lib.Tview.SetFocus(cfg.CloseFocus)
		},
	})
	lib.Winman.AddWindow(wm)
	lib.Winman.Center(wm)

	lib.Tview.UpdateDraw(func() {
		lib.Tview.SetFocus(cfg.Content)
	})

	return wm
}

func RemoveModal(modal *winman.WindowBase) {
	lib.Tview.UpdateDraw(func() {
		lib.Winman.RemoveWindow(modal)
		lib.Tview.SetFocus(nil)
	})
}
