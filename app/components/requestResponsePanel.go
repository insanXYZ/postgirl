package components

import (
	"fmt"
	"postgirl/app/lib"
	"postgirl/app/model"

	"github.com/rivo/tview"
)

type RequestResponsePanel struct {
	currentModel *model.Request

	// inputurl
	inputUrl *tview.Flex
	submit   chan bool

	//attribute
	attribute *tview.Flex

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
	ly.AddItem(requestResponsePanel.attribute, 13, 1, false)
	ly.AddItem(requestResponsePanel.response, 0, 1, false)

	requestResponsePanel.root = ly

	c.RequestResponsePanel = &requestResponsePanel
}

func (r *RequestResponsePanel) listenChan() {
	for {
		select {
		case <-r.submit:
			lib.Tview.Stop()
			fmt.Println("submit button")
			// reqConfig := model.Request{
			// 	Method: r.currentModel.Method,
			// 	Url:    r.currentModel.Url,
			// }

			// _ = request.NewRequest(&reqConfig)

		}
	}
}

func (r *RequestResponsePanel) Root() tview.Primitive {
	return r.root
}
