package tui

import (
	"budgettui/pkg/budget"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const Currency = "Currency"
const DefaultCurrency = "Default currency"
const Name = "Name"
const InitialAmount = "Initial amount"
const UseInBudget = "Use in budget"
const FirstDayOfMonth = "First day of month"
const Account = "Account"
const Amount = "Amount"
const Description = "Description"
const Category = "Category"

func RunApp() {
	data, err := budget.OpenFile()
	if err != nil {
		Init()
		data, _ = budget.OpenFile()
	}
	Home(data)
}

func LoadMenu(menu *tview.List, account *tview.Frame, app *tview.Application, pages *tview.Pages) {
	menu.Clear()
	data, _ := budget.OpenFile()
	accountList := GetAccountList(data, app, menu, pages)
	for i, accountOb := range data.Budgets[data.CurrentBudgetID].Accounts {
		var singleRune rune
		for _, char := range fmt.Sprintf("%v", i+1) {
			singleRune = char
		}
		accountToDisplay := accountList[accountOb.ID]
		menu.AddItem(accountOb.Name, fmt.Sprintf("%v %v", accountOb.Balance, accountOb.Currency), singleRune, func() {

			account.SetPrimitive(accountToDisplay)
		})
	}
	menu.
		AddItem("Budget settings", "", 'b', func() {
			app.Stop()
		}).
		AddItem("Change budget", "", 'c', func() {
			app.Stop()
		}).
		AddItem("Quit", "", 'q', func() {
			app.Stop()
		})
	menu.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRight {
			app.SetFocus(account)
			return nil
		}
		return event
	})

}

func GetMainView(helpInfo *tview.TextView, menu *tview.List, account *tview.Frame, budgetInfo *tview.TextView) *tview.Grid {
	mainView := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 0).
		AddItem(helpInfo, 2, 0, 1, 2, 0, 0, false).
		AddItem(menu, 1, 0, 1, 1, 0, 0, true).
		AddItem(account, 1, 1, 1, 1, 0, 0, false).
		AddItem(budgetInfo, 0, 0, 1, 2, 0, 0, false)

	return mainView
}

func SetInputs(app *tview.Application, pages *tview.Pages) {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyF1 {
			pages.ShowPage("help")
			return nil
		}
		if event.Key() == tcell.KeyCtrlB {
			pages.ShowPage("createBudget")
			return nil
		}
		if event.Key() == tcell.KeyCtrlA {
			pages.ShowPage("createAccount")
			return nil
		}
		return event
	})
}
