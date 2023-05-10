package tui

import (
	"budgettui/pkg/budget"
	"budgettui/pkg/helper"
	"fmt"
	"github.com/rivo/tview"
)

func RunApp() {

	found := budget.CheckSaveFile()
	if !found {
		Init()
	}
	Home()
}

func LoadAppElements(ctx budget.Context) {

	pages := ctx[helper.Pages].(*tview.Pages)
	detailsFrame := ctx[helper.DetailsFrame].(*tview.Frame)
	helpInfo := ctx[helper.HelpInfo].(*tview.TextView)
	//topBar := ctx[helper.TopBar].(*tview.Grid)

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

	CreateTopBar(ctx)

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
	helpInfo := ctx[helper.HelpInfo].(*tview.TextView)
	menu := ctx[helper.Menu].(*tview.List)
	detailsFrame := ctx[helper.DetailsFrame].(*tview.Frame)
	budgetInfo := ctx[helper.TopBar].(*tview.Grid)
	mainView := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 0).
		AddItem(helpInfo, 2, 0, 1, 2, 0, 0, false).
		AddItem(menu, 1, 0, 1, 1, 0, 0, true).
		AddItem(detailsFrame, 1, 1, 1, 1, 0, 0, false).
		AddItem(budgetInfo, 0, 0, 1, 2, 0, 0, false)

	return mainView
}

func CreateTopBar(ctx budget.Context) {
	data := ctx[helper.Data].(*budget.Data)
	pages := ctx[helper.Pages].(*tview.Pages)
	budgetName := tview.NewTextView().SetText(data.Budgets[data.CurrentBudgetID].Name)

	dateForm := CreateChangeDateForm(ctx)
	currentMonth := tview.NewButton(fmt.Sprintf("%v %v", ctx[helper.CurrentMonth].(*budget.YearMonth).Month, ctx[helper.CurrentMonth].(*budget.YearMonth).Year))
	currentMonth.SetSelectedFunc(func() {
		pages.AddPage("dateForm", tview.NewGrid().
			SetColumns(0, 30, 0).
			SetRows(0, 9, 0).
			AddItem(dateForm, 1, 1, 1, 1, 0, 0, true), true, false)
		pages.ShowPage("dateForm")
	})

	topBar := ctx[helper.TopBar].(*tview.Grid)
	topBar.SetBorder(true)
	topBar.SetRows(1).SetColumns(0, 0, 0)
	topBar.
		AddItem(budgetName, 0, 0, 1, 1, 0, 0, false).
		AddItem(currentMonth, 0, 2, 1, 1, 0, 0, false)
}

func CreateChangeDateForm(ctx budget.Context) *tview.Form {
	pages := ctx[helper.Pages].(*tview.Pages)
	form := tview.NewForm()
	datesList := ctx[helper.Data].(*budget.Data).GetMonths()
	var formOptions []string

	for _, m := range datesList {
		stringOption := fmt.Sprintf("%v, %v", m.Month, m.Year)
		formOptions = append(formOptions, stringOption)
	}
	form.AddDropDown("Date", formOptions, 0, nil)
	chosen, _ := form.GetFormItemByLabel("Date").(*tview.DropDown).GetCurrentOption()
	form.AddButton("save", func() {
		ctx[helper.CurrentMonth] = datesList[chosen]
		pages.HidePage("dateForm")

	})

	return form
}
