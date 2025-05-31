package components

import (
	"postgirl/model"

	"github.com/rivo/tview"
)

type RequestResponsePanel struct {
	currentRequest *model.Request

	// sync params
	chanToParams   chan string
	chanToInputUrl chan string

	// inputBar
	inputBar *InputBar
	send     chan bool

	// attribute
	attribute *Attribute

	// response
	response *Response

	// root
	root *tview.Flex
}

func NewRequestResponsePanel(req *model.Request) *tview.Flex {
	requestResponsePanel := RequestResponsePanel{
		currentRequest: req,
		chanToParams:   make(chan string),
		chanToInputUrl: make(chan string),
		send:           make(chan bool),
	}

	go requestResponsePanel.listenChan()

	requestResponsePanel.NewInputUrl()
	requestResponsePanel.NewAttribute()
	requestResponsePanel.NewResponse()

	ly := tview.NewFlex()
	ly.SetDirection(tview.FlexRow)
	ly.AddItem(requestResponsePanel.inputBar.Root, 3, 1, false)
	ly.AddItem(requestResponsePanel.attribute.Root, 16, 1, false)
	ly.AddItem(requestResponsePanel.response.Root, 0, 1, false)

	requestResponsePanel.root = ly

	return requestResponsePanel.root
}

func (r *RequestResponsePanel) Root() tview.Primitive {
	return r.root
}
