package components

import (
	"github.com/epiclabs-io/winman"
	"github.com/rivo/tview"
)

type Components struct {
	TviewApp  *tview.Application
	Winman    *winman.Manager
	Sidebar   *tview.Flex
	InputUrl  *tview.Flex
	Attribute *tview.Flex
	Response  *tview.Flex
	Layout    *tview.Flex
}

func NewComponents(app *tview.Application, winman *winman.Manager) *Components {
	cmp := &Components{
		TviewApp: app,
		Winman:   winman,
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
