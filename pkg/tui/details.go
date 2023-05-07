package tui

import (
	"budgettui/pkg/budget"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strconv"
	"time"
)

func GetDetailsFrames(ctx budget.Context) map[uint]*tview.Grid {
	d, _ := budget.LoadJSONData()

	//app := ctx[App].(*tview.Application)
	//mainMenu := ctx[Menu].(*tview.List)
	pages := ctx[Pages].(*tview.Pages)

	gridList := map[uint]*tview.Grid{}
	accounts := d.Budgets[d.CurrentBudgetID].Accounts
	transactions := tview.NewTable().SetBorders(true)

	for _, account := range accounts {
		//Buttons
		pageName := fmt.Sprintf("outcomeForm%v", account.Name)
		outcomeForm := GetOutcomeForm(account.ID, pageName, ctx)
		incomeForm := GetIncomeForm(account.ID, pageName, ctx)

		outcome := tview.NewButton("OUT").SetSelectedFunc(func() {
			pages.AddPage(pageName, tview.NewGrid().
				SetColumns(0, 58, 0).
				SetRows(0, 13, 0).
				AddItem(outcomeForm, 1, 1, 1, 1, 0, 0, true), true, false)
			pages.ShowPage(pageName)
		})
		income := tview.NewButton("IN").SetSelectedFunc(func() {
			pages.AddPage(pageName, tview.NewGrid().
				SetColumns(0, 58, 0).
				SetRows(0, 13, 0).
				AddItem(incomeForm, 1, 1, 1, 1, 0, 0, true), true, false)
			pages.ShowPage(pageName)
		})

		menu := tview.NewGrid().
			SetRows(1).
			SetColumns(0, 5, 1, 5, 5, 5, 0).
			AddItem(outcome, 0, 1, 1, 1, 0, 0, true).
			AddItem(income, 0, 3, 1, 1, 0, 0, false)

		outcome.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			//if event.Key() == tcell.KeyRight {
			//	app.SetFocus(income)
			//}
			//if event.Key() == tcell.KeyLeft {
			//	app.SetFocus(mainMenu)
			//}
			//if event.Key() == tcell.KeyDown {
			//	app.SetFocus(transactions)
			//}
			return event
		})
		income.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			//if event.Key() == tcell.KeyLeft {
			//	app.SetFocus(outcome)
			//}
			return event
		})

		acccountName := tview.NewTextView().SetText(account.Name)

		headers := tview.NewTable().SetBorders(true)
		headers.
			SetCell(0, 0, tview.NewTableCell(fmt.Sprintf("%-3.3s", "ID")).SetBackgroundColor(tcell.ColorDarkGrey)).
			SetCell(0, 1, tview.NewTableCell(fmt.Sprintf("%-30.30s", Description)).SetBackgroundColor(tcell.ColorDarkGrey)).
			SetCell(0, 2, tview.NewTableCell(fmt.Sprintf("%-15.15s", Amount)).SetBackgroundColor(tcell.ColorDarkGrey)).
			SetCell(0, 3, tview.NewTableCell(fmt.Sprintf("%-10.10s", Category)).SetBackgroundColor(tcell.ColorDarkGrey)).
			SetCell(0, 4, tview.NewTableCell(fmt.Sprintf("%-10.10s", "Date")).SetBackgroundColor(tcell.ColorDarkGrey))
		transactions.SetSelectable(true, false)
		for i, t := range d.Budgets[d.CurrentBudgetID].Accounts[account.ID].Transactions {
			transactions.
				SetCell(i, 0, tview.NewTableCell(fmt.Sprintf("%-3.3s", strconv.Itoa(int(t.ID))))).
				SetCell(i, 1, tview.NewTableCell(fmt.Sprintf("%-30.30s", t.Description))).
				SetCell(i, 2, tview.NewTableCell(fmt.Sprintf("%-15.15s", fmt.Sprintf("%.2f", t.Amount)))).
				SetCell(i, 3, tview.NewTableCell(fmt.Sprintf("%-10.10s", t.Category))).
				SetCell(i, 4, tview.NewTableCell(fmt.Sprintf("%-10.10s", t.Date.Format(time.DateOnly))))
		}
		transactions.ScrollToEnd()
		transactions.SetSelectedFunc(func(row, column int) {
			transactionForm := GetTransactionForm(int(account.ID), row, pageName, ctx)
			pages.AddPage(pageName, tview.NewGrid().
				SetColumns(0, 58, 0).
				SetRows(0, 13, 0).
				AddItem(transactionForm, 1, 1, 1, 1, 0, 0, true), true, false)
			pages.ShowPage(pageName)
		})
		accountGrid := tview.NewGrid().
			SetRows(1, 2, 3, 0).
			SetColumns(0, 0).
			AddItem(acccountName, 0, 0, 1, 7, 0, 0, false).
			AddItem(menu, 1, 0, 1, 2, 0, 0, true).
			AddItem(headers, 2, 0, 1, 4, 0, 0, true).
			AddItem(transactions, 3, 0, 1, 4, 0, 0, false)
		gridList[account.ID] = accountGrid
	}

	return gridList
}
