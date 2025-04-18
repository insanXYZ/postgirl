package components

import (
	"postgirl/app/internal/request"

	"github.com/rivo/tview"
)

type RequestResponsePanel struct {
	// inputurl
	inputUrl *tview.Flex
	method   string
	url      string
	submit   chan bool

	//attribute
	attribute *tview.Flex
	params    map[string]string
	headers   map[string]string
	bodyType  string
	body      string

	//response
	response *tview.Flex

	//root
	root *tview.Flex
}

func (c *Components) NewRequestResponsePanel() {
	requestResponsePanel := RequestResponsePanel{
		submit:  make(chan bool),
		params:  make(map[string]string),
		headers: make(map[string]string),
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
			reqConfig := request.RequestConfig{
				Method: r.method,
				Url:    r.url,
			}

			_ = request.NewRequest(&reqConfig)

		}
	}
}

func (r *RequestResponsePanel) Root() tview.Primitive {
	return r.root
}
