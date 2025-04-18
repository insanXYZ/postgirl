package components

import (
	"github.com/rivo/tview"
)

type Components struct {
	Sidebar *Sidebar

	//panel
	RequestResponsePanel     *RequestResponsePanel
	RequestResponsePanelChan chan string

	Layout *tview.Flex
}

func NewComponents() *Components {
	cmp := &Components{
		RequestResponsePanelChan: make(chan string),
	}

	go cmp.listenChangesRequestResponsePanel()

	cmp.NewSidebar()
	cmp.NewLayout()

	return cmp
}

func (c *Components) Root() tview.Primitive {
	return c.Layout
}
