package tui

import (
	"budgettui/pkg/budget"
	"budgettui/pkg/helper"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func LoadAppMenu(ctx budget.Context) {
	app := ctx[helper.App].(*tview.Application)
	detailsFrame := ctx[helper.DetailsFrame].(*tview.Frame)
	menu := ctx[helper.Menu].(*tview.List)
	currentItem := menu.GetCurrentItem()
	menu.Clear()

	data := ctx[helper.Data].(*budget.Data)

	accountList := GetDetailsFrames(ctx)

	budgetPage := tview.NewGrid().AddItem(tview.NewTextView().SetText("BUDGET"), 0, 0, 1, 1, 0, 0, false)
	var currentPrimitive tview.Primitive = budgetPage
	menu.
		//TODO Add budget page
		AddItem("Budget", "", '0', func() {
			detailsFrame.SetPrimitive(budgetPage)
		})

	for i, accountOb := range data.Budgets[data.CurrentBudgetID].Accounts {
		var singleRune rune
		for _, char := range fmt.Sprintf("%v", i+1) {
			singleRune = char
		}

		accountToDisplay := accountList[accountOb.ID]
		if i+1 == currentItem {
			currentPrimitive = accountToDisplay
		}
		menu.AddItem(accountOb.Name, fmt.Sprintf("%v %v", accountOb.Balance, accountOb.Currency), singleRune, func() {
			detailsFrame.SetPrimitive(accountToDisplay)
			app.SetFocus(accountToDisplay)
		})
	}
	menu.
		AddItem("Budget settings", "", 'b', func() {
			ShowPopup("Not ready", "alert", ctx)
		}).
		AddItem("Change budget", "", 'c', func() {
			app.Stop()
		}).
		AddItem("Quit", "", 'q', func() {
			ShowPopupQuit(helper.Alert, ctx)
		})
	menu.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		//if event.Key() == tcell.KeyRight {
		//	app.SetFocus(detailsFrame)
		//	return nil
		//}
		return event
	})

	menu.SetCurrentItem(currentItem)
	detailsFrame.SetPrimitive(currentPrimitive)
}
