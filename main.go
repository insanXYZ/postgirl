package main

import "postgirl/app"

func main() {
	app := app.NewApp()
	err := app.Run()

	if err != nil {
		panic(err.Error())
	}
}
