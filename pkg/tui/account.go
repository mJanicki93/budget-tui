package tui

import (
	"budgettui/pkg/budget"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strconv"
	"time"
)

func GetAccountList(d budget.Data, app *tview.Application, mainMenu *tview.List, pages *tview.Pages, accountFrame *tview.Frame) map[uint]*tview.Grid {
	gridList := map[uint]*tview.Grid{}
	accounts := d.Budgets[d.CurrentBudgetID].Accounts

	for _, account := range accounts {
		//Buttons
		pageName := fmt.Sprintf("outcomeForm%v", account.Name)
		outcomeForm := GetOutcomeForm(d, account.ID, pages, pageName, mainMenu, accountFrame, app)
		incomeForm := GetIncomeForm(d, account.ID, pages, pageName, mainMenu, accountFrame, app)

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
			if event.Key() == tcell.KeyRight {
				app.SetFocus(income)
			}
			if event.Key() == tcell.KeyLeft {
				app.SetFocus(mainMenu)
			}
			return event
		})
		income.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyLeft {
				app.SetFocus(outcome)
			}
			return event
		})

		acccountName := tview.NewTextView().SetText(account.Name)

		headers := tview.NewTable().SetBorders(true)
		headers.
			SetCell(0, 0, tview.NewTableCell(fmt.Sprintf("%-3.3s", "ID")).SetBackgroundColor(tcell.ColorDarkGrey)).
			SetCell(0, 1, tview.NewTableCell(fmt.Sprintf("%-20.20s", Description)).SetBackgroundColor(tcell.ColorDarkGrey)).
			SetCell(0, 2, tview.NewTableCell(fmt.Sprintf("%-15.15s", Amount)).SetBackgroundColor(tcell.ColorDarkGrey)).
			SetCell(0, 3, tview.NewTableCell(fmt.Sprintf("%-10.10s", Category)).SetBackgroundColor(tcell.ColorDarkGrey)).
			SetCell(0, 4, tview.NewTableCell(fmt.Sprintf("%-10.10s", "Date")).SetBackgroundColor(tcell.ColorDarkGrey))

		transactions := tview.NewTable().SetBorders(true)
		for i, t := range d.Budgets[d.CurrentBudgetID].Accounts[account.ID].Transactions {
			transactions.
				SetCell(i, 0, tview.NewTableCell(fmt.Sprintf("%-3.3s", strconv.Itoa(int(t.ID))))).
				SetCell(i, 1, tview.NewTableCell(fmt.Sprintf("%-20.20s", t.Description))).
				SetCell(i, 2, tview.NewTableCell(fmt.Sprintf("%-15.15s", fmt.Sprintf("%.2f", t.Amount)))).
				SetCell(i, 3, tview.NewTableCell(fmt.Sprintf("%-10.10s", t.Category))).
				SetCell(i, 4, tview.NewTableCell(fmt.Sprintf("%-10.10s", t.Date.Format(time.DateOnly))))
		}
		transactions.ScrollToEnd()
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

func NewCreateAccountFormInit() *tview.Form {
	form := tview.NewForm().
		AddInputField(Name, "", 20, nil, nil).
		AddDropDown(Currency, []string{"PLN", "EUR", "USD", "AUD"}, 0, nil).
		AddInputField(InitialAmount, "", 20, func(textToCheck string, lastChar rune) bool {
			floatValue, err := strconv.ParseFloat(textToCheck, 64)
			if err != nil {
				return false
			}
			if floatValue < 0 {
				return false
			}
			return true
		}, nil).
		AddCheckbox(UseInBudget, true, nil)

	form.SetBorder(true).SetTitle("New Account").SetTitleAlign(tview.AlignLeft)

	return form
}

func NewCreateAccountForm(data budget.Data, pages *tview.Pages, menu *tview.List, account *tview.Frame, app *tview.Application) *tview.Form {
	form := tview.NewForm().
		AddInputField(Name, "", 20, nil, nil).
		AddDropDown(Currency, []string{"PLN", "EUR", "USD", "AUD"}, 0, nil).
		AddInputField(InitialAmount, "", 20, func(textToCheck string, lastChar rune) bool {
			floatValue, err := strconv.ParseFloat(textToCheck, 64)
			if err != nil {
				return false
			}
			if floatValue < 0 {
				return false
			}
			return true
		}, nil).
		AddCheckbox(UseInBudget, true, nil)

	form.SetBorder(true).SetTitle("New Account").SetTitleAlign(tview.AlignLeft)

	form.AddButton("Save", func() {
		//Get form values
		accountName := form.GetFormItemByLabel(Name).(*tview.InputField).GetText()
		_, currency := form.GetFormItemByLabel(Currency).(*tview.DropDown).GetCurrentOption()
		initialAmount, _ := strconv.ParseFloat(form.GetFormItemByLabel(InitialAmount).(*tview.InputField).GetText(), 64)
		useInBudget := form.GetFormItemByLabel(UseInBudget).(*tview.Checkbox).IsChecked()

		//Add account entity
		data.Budgets[data.CurrentBudgetID].Accounts = append(data.Budgets[data.CurrentBudgetID].Accounts, budget.Account{
			ID:          uint(len(data.Budgets[data.CurrentBudgetID].Accounts)),
			Name:        accountName,
			Currency:    currency,
			Balance:     initialAmount,
			UseInBudget: useInBudget,
		})

		//Actions
		_ = data.SaveFile()
		LoadMenu(menu, account, app, pages)
		pages.HidePage("createAccount")
		pages.ShowPage("main")
	}).
		AddButton("Quit", func() {
			pages.HidePage("createAccount")
			pages.ShowPage("main")
		})

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			pages.HidePage("createAccount")
			pages.ShowPage("main")
		}
		return event
	})

	return form
}
