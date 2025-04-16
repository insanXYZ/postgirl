package components

import "github.com/rivo/tview"

func (c *Components) NewAttribute() {
	paramsButton := tview.NewButton("Params")
	headersButton := tview.NewButton("Headers")
	bodyButton := tview.NewButton("Body")

	flexButton := tview.NewFlex()
	flexButton.AddItem(paramsButton, 10, 1, false)
	flexButton.AddItem(tview.NewBox(), 1, 1, false)
	flexButton.AddItem(headersButton, 10, 1, false)
	flexButton.AddItem(tview.NewBox(), 1, 1, false)
	flexButton.AddItem(bodyButton, 10, 1, false)
	flexButton.AddItem(tview.NewBox(), 1, 1, false)
	flexButton.SetBorder(true)

	flexAttribute := tview.NewFlex()
	flexAttribute.SetDirection(tview.FlexRow)
	flexAttribute.AddItem(flexButton, 3, 1, false)
	flexAttribute.AddItem(tview.NewBox().SetBorder(true), 10, 1, false)

	removedBodyDropdown := func() {
		if flexButton.GetItemCount() == 7 {
			go c.TviewApp.QueueUpdateDraw(func() {
				flexButton.RemoveItem(flexButton.GetItem(flexButton.GetItemCount() - 1))
			})
		}
	}

	bodyButton.SetSelectedFunc(func() {
		go c.TviewApp.QueueUpdateDraw(func() {
			dropdownBody := tview.NewDropDown()

			dropdownBody.SetOptions(
				[]string{"none",
					"form-data",
					"x-www-form-urlencoded",
					"json",
					"xml"},

				func(text string, index int) {

				},
			)

			flexButton.AddItem(dropdownBody, 15, 1, false)
		})
	})

	paramsButton.SetSelectedFunc(func() {
		removedBodyDropdown()
	})

	headersButton.SetSelectedFunc(func() {
		removedBodyDropdown()
	})

	c.Attribute = flexAttribute
}
