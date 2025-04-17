package components

import (
	"github.com/rivo/tview"
)

type Components struct {
	Sidebar     *Sidebar
	EditorPanel *EditorPanel
	Layout      *tview.Flex
}

func NewComponents() *Components {
	cmp := &Components{}
	cmp.NewSidebar()
	cmp.NewEditorPanel()
	cmp.NewLayout()

	return cmp
}

func (c *Components) Root() tview.Primitive {
	return c.Layout
}
