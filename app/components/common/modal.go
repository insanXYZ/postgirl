package common

import (
	"postgirl/app/color"
	"postgirl/app/lib"

	"github.com/epiclabs-io/winman"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ModalConfig struct {
	Content       tview.Primitive
	CloseFocus    tview.Primitive
	CloseButton   bool
	Center        bool
	X, Y          int
	Width, Height int
	Title         string
	TitleAlign    int
	BorderColor   tcell.Color
}

func ShowModal(cfg *ModalConfig) *winman.WindowBase {
	wm := lib.Winman.NewWindow().Show()
	wm.SetTitleAlign(cfg.TitleAlign)
	wm.SetTitle(cfg.Title)
	wm.SetRoot(cfg.Content)
	wm.SetModal(true)
	wm.SetBorder(true)
	wm.SetBackgroundColor(color.BACKGROUND)
	wm.SetTitleColor(color.BORDER)
	wm.SetBorderColor(cfg.BorderColor)
	wm.SetRect(cfg.X, cfg.Y, cfg.Width, cfg.Height)

	lib.Winman.AddWindow(wm)

	if cfg.CloseButton {
		wm.AddButton(&winman.Button{
			Symbol: 'X',
			OnClick: func() {
				lib.Winman.RemoveWindow(wm)
				lib.Tview.SetFocus(cfg.CloseFocus)
			},
		})
	}

	if cfg.Center {
		lib.Winman.Center(wm)
	}

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
