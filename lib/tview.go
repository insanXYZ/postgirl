package lib

import "github.com/rivo/tview"

var Tview *tviewApp

func init() {
	Tview = newTviewApp()
}

type tviewApp struct {
	*tview.Application
}

func newTviewApp() *tviewApp {
	return &tviewApp{
		Application: tview.NewApplication(),
	}
}

func (t *tviewApp) UpdateDraw(f func()) {
	go t.QueueUpdateDraw(f)
}
