package tui

import (
	"budgettui/pkg/budget"
	"github.com/rivo/tview"
)

func RunApp() {

	_, err := budget.LoadJSONData()
	if err != nil {
		Init()
	}
	Home()
}

func LoadAppElements(ctx budget.Context) {
	data, _ := budget.LoadJSONData()

	pages := ctx[Pages].(*tview.Pages)
	detailsFrame := ctx[DetailsFrame].(*tview.Frame)
	helpInfo := ctx[HelpInfo].(*tview.TextView)
	budgetInfo := ctx[BudegtInfo].(*tview.TextView)

	LoadAppMenu(ctx)

	helpFrame := GetHelpWindow(ctx)
	newBudget := NewCreateBudgetForm(ctx)
	newAccount := NewCreateAccountForm(ctx)
	transferForm := GetTransferForm(ctx)
	quickOutcomeForm := GetQuickTransactionForm(false, ctx)
	quickIncomeForm := GetQuickTransactionForm(true, ctx)

	helpInfo.SetBorder(true)
	helpInfo.
		SetText(" Press F1 for helpFrame, press Ctrl-C to exit")

	budgetInfo.SetBorder(true)
	budgetInfo.
		SetText(data.Budgets[data.CurrentBudgetID].Name)

	detailsFrame.SetBorder(true)

	mainGrid := GetMainView(ctx)

	pages.AddAndSwitchToPage("main", mainGrid, true).
		AddPage("helpFrame", tview.NewGrid().
			SetColumns(0, 64, 0).
			SetRows(0, 22, 0).
			AddItem(helpFrame, 1, 1, 1, 1, 0, 0, true), true, false).
		AddPage("createBudget", tview.NewGrid().
			SetColumns(0, 58, 0).
			SetRows(0, 12, 0).
			AddItem(newBudget, 1, 1, 1, 1, 0, 0, true), true, false).
		AddPage("createAccount", tview.NewGrid().
			SetColumns(0, 58, 0).
			SetRows(0, 13, 0).
			AddItem(newAccount, 1, 1, 1, 1, 0, 0, true), true, false).
		AddPage("transferForm", tview.NewGrid().
			SetColumns(0, 58, 0).
			SetRows(0, 15, 0).
			AddItem(transferForm, 1, 1, 1, 1, 0, 0, true), true, false).
		AddPage("quickOutcome", tview.NewGrid().
			SetColumns(0, 30, 0).
			SetRows(0, 9, 0).
			AddItem(quickOutcomeForm, 1, 1, 1, 1, 0, 0, true), true, false).
		AddPage("quickIncome", tview.NewGrid().
			SetColumns(0, 30, 0).
			SetRows(0, 9, 0).
			AddItem(quickIncomeForm, 1, 1, 1, 1, 0, 0, true), true, false)

}

func GetMainView(ctx budget.Context) *tview.Grid {
	helpInfo := ctx[HelpInfo].(*tview.TextView)
	menu := ctx[Menu].(*tview.List)
	detailsFrame := ctx[DetailsFrame].(*tview.Frame)
	budgetInfo := ctx[BudegtInfo].(*tview.TextView)
	mainView := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 0).
		AddItem(helpInfo, 2, 0, 1, 2, 0, 0, false).
		AddItem(menu, 1, 0, 1, 1, 0, 0, true).
		AddItem(detailsFrame, 1, 1, 1, 1, 0, 0, false).
		AddItem(budgetInfo, 0, 0, 1, 2, 0, 0, false)

	return mainView
}
