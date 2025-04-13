package app

import (
	"postgirl/app/components"

	"github.com/rivo/tview"
)

type App struct {
	tviewApp   *tview.Application
	components *components.Components
}

func NewApp() *App {
	app := &App{}
	app.tviewApp = tview.NewApplication()
	app.components = components.NewComponents(app.tviewApp)
	return app
}

func (a *App) Run() error {
	return a.tviewApp.SetRoot(a.components.Root(), true).EnableMouse(true).Run()
}
