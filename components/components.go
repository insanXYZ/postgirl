package components

import (
	"github.com/rivo/tview"
)

var switchRequestResponsePanelChan chan string

func init() {
	switchRequestResponsePanelChan = make(chan string)
}

type Components struct {
	Sidebar              *Sidebar
	RequestResponsePanel *RequestResponsePanel
	Layout               *tview.Flex
}

func NewComponents() *Components {
	cmp := &Components{}

	go cmp.listenChan()

	cmp.Sidebar = NewSidebar()
	cmp.NewLayout()

	return cmp
}

func (c *Components) listenChan() {
	for label := range switchRequestResponsePanelChan {
		c.ChangePanel(label)
	}
}

func (c *Components) Root() tview.Primitive {
	return c.Layout
}
