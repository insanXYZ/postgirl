package components

import (
	"fmt"
	"postgirl/app/lib"

	"github.com/rivo/tview"
)

type EditorPanel struct {
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

func (c *Components) NewEditorPanel() {
	editorPanel := EditorPanel{
		submit:  make(chan bool),
		params:  make(map[string]string),
		headers: make(map[string]string),
	}

	go editorPanel.listenChan()

	editorPanel.NewInputUrl()
	editorPanel.NewAttribute()
	editorPanel.NewResponse()

	ly := tview.NewFlex()
	ly.SetDirection(tview.FlexRow)
	ly.AddItem(editorPanel.inputUrl, 3, 1, false)
	ly.AddItem(editorPanel.attribute, 13, 1, false)
	ly.AddItem(editorPanel.response, 0, 1, false)

	editorPanel.root = ly

	c.EditorPanel = &editorPanel
}

func (e *EditorPanel) listenChan() {
	for {
		select {
		case <-e.submit:
			lib.Tview.Stop()
			fmt.Println(e.method, e.url)
		}
	}
}

func (e *EditorPanel) Root() tview.Primitive {
	return e.root
}
