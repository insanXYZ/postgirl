package common

import (
	"postgirl/color"

	"github.com/rivo/tview"
)

type List struct {
	*tview.List
}

type ListConfig struct {
	Border          bool
	SetSelectedFunc func(i int, s string, isCheckbox bool, checked bool)
}

func CreateList(cfg *ListConfig) *List {
	list := List{}

	l := tview.NewList()
	l.SetBorder(cfg.Border)
	l.SetBorderColor(color.BORDER)
	l.SetBackgroundColor(color.BACKGROUND)
	list.List = l

	l.SetSelectedFunc(func(i int, s1, s2 string, r rune) {
		var checked bool
		isCheckbox := list.isCheckbox(s1)
		runeLabel := []rune(s1)
		label := s1

		if isCheckbox {
			label = string(runeLabel[4:])
		}

		if isCheckbox {
			switch runeLabel[1] {
			case 'X':
				runeLabel[1] = ' '
				checked = false
			case ' ':
				runeLabel[1] = 'X'
				checked = true
			}

			l.SetItemText(i, string(runeLabel), "")
		}

		if cfg.SetSelectedFunc != nil {
			cfg.SetSelectedFunc(i, label, isCheckbox, checked)
		}

	})

	return &list
}

func (l *List) GetRoot() *tview.List {
	return l.List
}

func (l *List) RemoveAll() {
	for i := l.List.GetItemCount(); i >= 0; i-- {
		l.List.RemoveItem(i)
	}
}

func (l *List) RemoveItem(i int) {
	l.List.RemoveItem(i)
}

func (l *List) isCheckbox(s string) bool {

	return s[0] == '(' && s[2:4] == ") " && len(s) > 4
}

func (l *List) AddItem(label string, checkBox bool) {

	if checkBox {
		label = "( ) " + label
	}

	l.List.AddItem(label, "", 0, nil)
}
