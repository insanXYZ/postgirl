package components

import (
	"postgirl/app/model"

	"github.com/rivo/tview"
)

type RequestResponsePanel struct {
	currentModel *model.Request

	// inputurl
	inputUrl *tview.Flex
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
	ly.AddItem(requestResponsePanel.inputUrl, 3, 1, false)
	ly.AddItem(requestResponsePanel.attribute.Root, 16, 1, false)
	ly.AddItem(requestResponsePanel.response, 0, 1, false)

	requestResponsePanel.root = ly

	c.RequestResponsePanel = &requestResponsePanel
}

func (r *RequestResponsePanel) listenChan() {
	for {
		select {
		case <-r.submit:
			// process params json and headers json to map

			req := model.Request{
				Url:       r.currentModel.Url,
				Method:    r.currentModel.Method,
				Attribute: model.Attribute{},
			}
		}
	}
}

func (r *RequestResponsePanel) Root() tview.Primitive {
	return r.root
}
