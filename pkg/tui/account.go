package tui

import (
	"budgettui/pkg/budget"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strconv"
)

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

func NewCreateAccountForm(ctx budget.Context) *tview.Form {
	data, _ := budget.LoadJSONData()
	pages := ctx[Pages].(*tview.Pages)

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

		//Add detailsFrame entity
		data.Budgets[data.CurrentBudgetID].Accounts = append(data.Budgets[data.CurrentBudgetID].Accounts, budget.Account{
			ID:          uint(len(data.Budgets[data.CurrentBudgetID].Accounts)),
			Name:        accountName,
			Currency:    currency,
			Balance:     initialAmount,
			UseInBudget: useInBudget,
		})

		//Actions
		_ = data.SaveFile()
		LoadAppData(ctx)
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
