package components

import "github.com/rivo/tview"

type Components struct {
	Sidebar     *tview.List
	EditorPanel *EditorPanel
	Layout      *tview.Flex
}

func NewComponents() *Components {
	cmp := &Components{
		Sidebar:     NewSidebar(),
		EditorPanel: NewEditorPanel(),
	}

	layout := NewLayout(cmp.Sidebar, NewEditorPanel())
	cmp.Layout = layout

	return cmp
}

func (c *Components) Root() tview.Primitive {
	return c.Layout
}
