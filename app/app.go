package app

import (
	"postgirl/app/components"
	"postgirl/app/lib"
)

type App struct {
	components *components.Components
}

func NewApp() *App {
	app := &App{}
	app.components = components.NewComponents()
	return app
}

func (a *App) Run() error {
	lib.Winman.NewWindow().SetRoot(a.components.Root()).Maximize().SetBorder(false).Show()
	return lib.Tview.SetRoot(lib.Winman, true).EnableMouse(true).Run()
}
