package components

import "github.com/rivo/tview"

func NewLayout(sidebar *tview.List, editorPanel *EditorPanel) *tview.Flex {
	flex := tview.NewFlex()
	flex.AddItem(sidebar, 0, 1, false)
	flex.AddItem(editorPanel.Root(), 0, 1, false)

	return flex
}
