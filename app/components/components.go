package components

import "github.com/rivo/tview"

type Components struct {
	TviewApp  *tview.Application
	Sidebar   *tview.List
	InputUrl  *tview.Flex
	Attribute *tview.Flex
	Response  *tview.Flex
	Layout    *tview.Flex
}

func NewComponents(app *tview.Application) *Components {
	cmp := &Components{
		TviewApp: app,
	}
	cmp.NewSidebar()
	cmp.NewInputUrl()
	cmp.NewAttribute()
	cmp.NewResponse()
	cmp.NewLayout()

	return cmp
}

func (c *Components) Root() tview.Primitive {
	return c.Layout
}
