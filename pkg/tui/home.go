package tui

import (
	"budgettui/pkg/budget"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Home() {
	app := tview.NewApplication()

	debug := true

	data := budget.OpenFile()

	accountList := GetAccountList(data)
	debugWindow := tview.NewTextView().SetText(fmt.Sprintf("%v", accountList))

	pages := tview.NewPages()
	menu := tview.NewList()
	budgetInfo := tview.NewTextView()
	account := tview.NewFrame(accountList[0])
	helpInfo := tview.NewTextView()

	helpInfo.SetBorder(true)
	helpInfo.
		SetText(" Press F1 for help, press Ctrl-C to exit")

	budgetInfo.SetBorder(true)
	budgetInfo.
		SetText(data.Budgets[data.CurrentBudgetID].Name)

	account.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyESC {
			app.SetFocus(menu)
			return nil
		}
		return event
	})
	account.SetBorder(true)
	for i, accountOb := range data.Budgets[data.CurrentBudgetID].Accounts {
		var singleRune rune
		for _, char := range fmt.Sprintf("%v", i) {
			singleRune = char
		}
		accountToDisplay := accountList[accountOb.ID]
		menu.AddItem(accountOb.Name, accountOb.Currency, singleRune, func() {

			account.SetPrimitive(accountToDisplay)

			debugWindow.SetText(fmt.Sprintf("i: %v, account: %v", i, accountOb))
			app.SetFocus(account)
		})
	}
	menu.
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})

	mainView := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 0).
		AddItem(helpInfo, 2, 0, 1, 2, 0, 0, false).
		AddItem(menu, 1, 0, 1, 1, 0, 0, true).
		AddItem(account, 1, 1, 1, 1, 0, 0, false).
		AddItem(budgetInfo, 0, 0, 1, 2, 0, 0, false)

	if debug {
		mainView.SetRows(3, 0, 3, 10)
		mainView.AddItem(debugWindow, 3, 0, 1, 2, 0, 0, false)
	}
	pages.AddAndSwitchToPage("main", mainView, true)

	if err := app.SetRoot(pages,
		true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
