package components

import "github.com/rivo/tview"

type EditorPanel struct {
	inputUrl  *tview.Flex
	attribute *tview.Flex
	response  *tview.Flex
}

func NewEditorPanel() *EditorPanel {
	return &EditorPanel{
		inputUrl:  NewInputUrl(),
		attribute: NewInputUrl(),
		response:  NewResponse(),
	}
}

func (e *EditorPanel) Root() *tview.Flex {
	flex := tview.NewFlex()
	flex.AddItem(e.inputUrl, 0, 1, false)
	flex.AddItem(e.attribute, 0, 1, false)
	flex.AddItem(e.response, 0, 1, false)

	return flex
}
