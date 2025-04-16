package app

import (
	"postgirl/app/components"

	"github.com/epiclabs-io/winman"
	"github.com/rivo/tview"
)

type App struct {
	tviewApp   *tview.Application
	winman     *winman.Manager
	components *components.Components
}

func NewApp() *App {
	app := &App{}
	app.tviewApp = tview.NewApplication()
	app.winman = winman.NewWindowManager()
	app.components = components.NewComponents(app.tviewApp, app.winman)
	return app
}

func (a *App) Run() error {
	a.winman.NewWindow().SetRoot(a.components.Root()).Maximize().SetBorder(false).Show()
	return a.tviewApp.SetRoot(a.winman, true).EnableMouse(true).Run()
}
