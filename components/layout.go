package components

import (
	"postgirl/color"
	"postgirl/components/common"
	"postgirl/internal/cache"
	"postgirl/lib"

	"github.com/rivo/tview"
)

func (c *Components) NewLayout() {
	var requestResponsePanel tview.Primitive

	if c.RequestResponsePanel != nil {
		requestResponsePanel = c.RequestResponsePanel.Root()
	} else {
		requestResponsePanel = common.CreateEmptyBox()
	}

	layoutFlex := tview.NewFlex()
	layoutFlex.AddItem(c.Sidebar.Root(), 30, 1, false)
	layoutFlex.AddItem(requestResponsePanel, 0, 1, false)
	layoutFlex.SetBackgroundColor(color.BACKGROUND)

	c.Layout = layoutFlex
}

func (c *Components) ChangePanel(label string) {
	panel := cache.CacheRequests.GetPanel(label)

	if panel == nil {
		cache.CacheRequests.SetPanel(label, NewRequestResponsePanel(cache.CacheRequests.GetRequest(label)))
		panel = cache.CacheRequests.GetPanel(label)
	}

	if c.Layout.GetItemCount() == 2 {
		c.Layout.RemoveItem(c.Layout.GetItem(c.Layout.GetItemCount() - 1))
	}

	lib.Tview.UpdateDraw(func() {
		c.Layout.AddItem(panel, 0, 1, false)
	})
}
