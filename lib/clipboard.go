package lib

import "golang.design/x/clipboard"

func NewClipboard() error {
	return clipboard.Init()
}
