package main

import (
	"fmt"
	"os"
	"postgirl/components"
	"postgirl/lib"
	"postgirl/model"
)

func main() {

	throwError := func(e error) {
		fmt.Printf("%v\n%v", model.ErrStartApp, "error detail :"+e.Error())
		os.Exit(0)
	}

	err := lib.NewClipboard()
	if err != nil {
		throwError(err)
	}

	components := components.NewComponents()

	lib.Winman.NewWindow().SetRoot(components.Root()).Maximize().SetBorder(false).Show()
	err = lib.Tview.SetRoot(lib.Winman, true).EnableMouse(true).Run()

	if err != nil {
		throwError(err)
	}

}
