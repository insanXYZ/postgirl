package main

import (
	"fmt"
	"postgirl/components"
	"postgirl/lib"
	"postgirl/model"
)

func main() {

	err := lib.NewClipboard()
	if err != nil {
		fmt.Printf("%v\n%v", model.ErrStartApp, "error detail :"+err.Error())
	}

	components := components.NewComponents()

	lib.Winman.NewWindow().SetRoot(components.Root()).Maximize().SetBorder(false).Show()
	err = lib.Tview.SetRoot(lib.Winman, true).EnableMouse(true).Run()

	if err != nil {
		fmt.Printf("%v\n%v", model.ErrStartApp, "error detail :"+err.Error())
	}
}
