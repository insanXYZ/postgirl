package components

import (
	"postgirl/app/model"

	"github.com/rivo/tview"
)

type RequestResponsePanel struct {
	currentModel *model.Request

	// inputBar
	inputBar *InputBar
	submit   chan bool

	//attribute
	attribute *Attribute

	//response
	response *tview.Flex

	//root
	root *tview.Flex
}

func (c *Components) NewRequestResponsePanel(req *model.Request) {
	requestResponsePanel := RequestResponsePanel{
		currentModel: req,
		submit:       make(chan bool),
	}

	go requestResponsePanel.listenChan()

	requestResponsePanel.NewInputUrl()
	requestResponsePanel.NewAttribute()
	requestResponsePanel.NewResponse()

	ly := tview.NewFlex()
	ly.SetDirection(tview.FlexRow)
	ly.AddItem(requestResponsePanel.inputBar.Root, 3, 1, false)
	ly.AddItem(requestResponsePanel.attribute.Root, 16, 1, false)
	ly.AddItem(requestResponsePanel.response, 0, 1, false)

	requestResponsePanel.root = ly

	c.RequestResponsePanel = &requestResponsePanel
}

func (r *RequestResponsePanel) Root() tview.Primitive {
	return r.root
}
