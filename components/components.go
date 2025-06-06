package components

import (
	"postgirl/components/common"
	"postgirl/internal/cache"
	"postgirl/model"

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

func SaveCache() {
	err := cache.CacheRequests.Save()
	if err != nil {
		common.ShowNotification(&common.NotificationConfig{
			Message: model.ErrSaveCache,
		})
	}
}
