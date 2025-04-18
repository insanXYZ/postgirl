package components

import (
	"github.com/rivo/tview"
)

type Components struct {
	Sidebar              *Sidebar
	RequestResponsePanel *RequestResponsePanel
	Layout               *tview.Flex
}

func NewComponents() *Components {
	cmp := &Components{}
	cmp.NewSidebar()
	cmp.NewRequestResponsePanel()
	cmp.NewLayout()

	return cmp
}

func (c *Components) Root() tview.Primitive {
	return c.Layout
}
