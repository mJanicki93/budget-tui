package tui

import (
	"budgettui/pkg/budget"
	"github.com/rivo/tview"
)

func GetAccountList(d budget.Data) map[uint]*tview.Grid {
	gridList := map[uint]*tview.Grid{}
	accounts := d.Budgets[d.CurrentBudgetID].Accounts

	for _, account := range accounts {
		acccountName := tview.NewTextView().SetText(account.Name)
		transactions := tview.NewTextView().SetText("Transacions")

		accountGrid := tview.NewGrid().
			SetRows(3, 0).
			SetColumns(60, 0).
			AddItem(acccountName, 0, 1, 1, 2, 0, 0, false).
			AddItem(transactions, 1, 1, 1, 1, 0, 0, false)
		gridList[account.ID] = accountGrid
	}

	return gridList
}
