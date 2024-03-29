package tui

import (
	"budgettui/pkg/budget"
	"budgettui/pkg/helper"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strconv"
	"time"
)

func GetDetailsFrames(ctx budget.Context) map[uint]*tview.Grid {
	d := ctx[helper.Data].(*budget.Data)

	app := ctx[helper.App].(*tview.Application)
	mainMenu := ctx[helper.Menu].(*tview.List)
	pages := ctx[helper.Pages].(*tview.Pages)

	gridList := map[uint]*tview.Grid{}
	accounts := d.Budgets[d.CurrentBudgetID].Accounts

	for _, account := range accounts {
		//Buttons
		outcome := tview.NewButton("OUT")
		income := tview.NewButton("IN")
		headers := tview.NewTable().SetBorders(false)
		outcome.SetBackgroundColor(tcell.ColorDarkRed)
		outcome.SetBackgroundColorActivated(tcell.ColorRed).SetTitleColor(tcell.ColorWhite)
		income.SetBackgroundColor(tcell.ColorDarkGreen)
		income.SetBackgroundColorActivated(tcell.ColorGreen).SetTitleColor(tcell.ColorWhite)

		pageName := fmt.Sprintf("outcomeForm%v", account.Name)
		outcomeForm := GetNewTransactionForm(account.ID, pageName, false, ctx)
		incomeForm := GetNewTransactionForm(account.ID, pageName, true, ctx)
		transactions := tview.NewTable().SetBorders(false)
		transactions.SetTitle(fmt.Sprintf("%v", account.ID))

		outcome.SetBackgroundColor(tcell.ColorDarkRed).SetTitleColor(tcell.ColorWhite)
		outcome.SetBackgroundColorActivated(tcell.ColorRed)
		income.SetBackgroundColor(tcell.ColorDarkGreen).SetTitleColor(tcell.ColorWhite)
		income.SetBackgroundColorActivated(tcell.ColorGreen)

		outcome.SetSelectedFunc(func() {
			pages.AddPage(pageName, tview.NewGrid().
				SetColumns(0, 58, 0).
				SetRows(0, 13, 0).
				AddItem(outcomeForm, 1, 1, 1, 1, 0, 0, true), true, false)
			pages.ShowPage(pageName)
		})
		income.SetSelectedFunc(func() {
			pages.AddPage(pageName, tview.NewGrid().
				SetColumns(0, 58, 0).
				SetRows(0, 13, 0).
				AddItem(incomeForm, 1, 1, 1, 1, 0, 0, true), true, false)
			pages.ShowPage(pageName)
		})

		menu := tview.NewGrid().
			SetRows(1).
			SetColumns(5, 1, 5, 5, 5, 0).
			AddItem(outcome, 0, 0, 1, 1, 0, 0, true).
			AddItem(income, 0, 2, 1, 1, 0, 0, false)

		outcome.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyTab {
				app.SetFocus(income)
			}
			if event.Key() == tcell.KeyEscape {
				app.SetFocus(mainMenu)
			}
			return event
		})
		income.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyTab {
				app.SetFocus(transactions)
				headers.SetBackgroundColor(tcell.ColorDarkBlue)
			}
			if event.Key() == tcell.KeyEscape {
				app.SetFocus(mainMenu)
			}
			return event
		})

		accountName := tview.NewTextView().SetText(account.Name)
		headers.SetBackgroundColor(tcell.ColorDarkGrey)
		headers.
			SetCell(0, 0, tview.NewTableCell(fmt.Sprintf("%-3.3s", "ID"))).
			SetCell(0, 1, tview.NewTableCell(fmt.Sprintf("%-30.30s", helper.Description))).
			SetCell(0, 2, tview.NewTableCell(fmt.Sprintf("%-15.15s", helper.Amount))).
			SetCell(0, 3, tview.NewTableCell(fmt.Sprintf("%-10.10s", helper.Category))).
			SetCell(0, 4, tview.NewTableCell(fmt.Sprintf("%-10.10s", "Date")))
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
			transactionsID, _ := strconv.Atoi(transactions.GetTitle())
			transactionForm := GetTransactionForm(uint(transactionsID), uint(row), pageName, ctx)
			pages.AddPage(pageName, tview.NewGrid().
				SetColumns(0, 58, 0).
				SetRows(0, 13, 0).
				AddItem(transactionForm, 1, 1, 1, 1, 0, 0, true), true, false)
			pages.ShowPage(pageName)
		})
		transactions.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyTab {
				app.SetFocus(outcome)
				headers.SetBackgroundColor(tcell.ColorDarkGrey)
			}
			if event.Key() == tcell.KeyEscape {
				app.SetFocus(mainMenu)
			}
			return event
		})
		accountGrid := tview.NewGrid().
			SetRows(1, 1, 1, 1, 0).
			SetColumns(0, 0).
			AddItem(accountName, 0, 0, 1, 7, 0, 0, false).
			AddItem(menu, 1, 0, 1, 2, 0, 0, true).
			AddItem(headers, 3, 0, 1, 4, 0, 0, true).
			AddItem(transactions, 4, 0, 1, 4, 0, 0, false)
		gridList[account.ID] = accountGrid
	}

	return gridList
}
