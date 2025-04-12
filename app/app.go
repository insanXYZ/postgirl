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
	return &App{
		tviewApp:   tview.NewApplication(),
		components: components.NewComponents(),
	}
}

func (a *App) Run() error {
	return a.tviewApp.SetRoot(a.components.Root(), true).Run()
}
