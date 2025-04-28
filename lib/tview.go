package lib

import (
	"github.com/rivo/tview"
)

var Tview *tviewApp

func init() {
	Tview = newTviewApp()
}

type tviewApp struct {
	*tview.Application
}

func newTviewApp() *tviewApp {
	app := tview.NewApplication()

	return &tviewApp{
		Application: app,
	}
}

func (t *tviewApp) UpdateDraw(f func()) {
	go t.QueueUpdateDraw(f)
}
