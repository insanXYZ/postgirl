package lib

import "github.com/epiclabs-io/winman"

var Winman *winman.Manager

func init() {
	Winman = winman.NewWindowManager()
}
