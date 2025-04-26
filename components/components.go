package components

import (
	"postgirl/internal/cache"

	"github.com/rivo/tview"
)

type Components struct {
	Sidebar *Sidebar

	//panel
	RequestResponsePanel           *RequestResponsePanel
	SwitchRequestResponsePanelChan chan string

	Layout *tview.Flex
}

func NewComponents() *Components {
	cmp := &Components{
		SwitchRequestResponsePanelChan: make(chan string),
	}

	go cmp.listenChan()

	cmp.NewSidebar()
	cmp.NewLayout()

	return cmp
}

func (c *Components) listenChan() {
	for {
		select {
		case label := <-c.SwitchRequestResponsePanelChan:
			cache := cache.CacheRequests.GetRequest(label)
			c.ChangePanel(cache)
		}
	}
}

func (c *Components) Root() tview.Primitive {
	return c.Layout
}
