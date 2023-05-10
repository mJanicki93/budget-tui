package tui

import (
	"budgettui/pkg/budget"
	"budgettui/pkg/helper"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strconv"
)

func NewCreateAccountFormInit() *tview.Form {
	form := tview.NewForm().
		AddInputField(helper.Name, "", 20, nil, nil).
		AddDropDown(helper.Currency, []string{"PLN", "EUR", "USD", "AUD"}, 0, nil).
		AddInputField(helper.InitialAmount, "", 20, func(textToCheck string, lastChar rune) bool {
			floatValue, err := strconv.ParseFloat(textToCheck, 64)
			if err != nil {
				return false
			}
			if floatValue < 0 {
				return false
			}
			return true
		}, nil).
		AddCheckbox(helper.UseInBudget, true, nil)

	form.SetBorder(true).SetTitle("New Account").SetTitleAlign(tview.AlignLeft)

	return form
}

func NewCreateAccountForm(ctx budget.Context) *tview.Form {
	data := ctx[helper.Data].(*budget.Data)
	pages := ctx[helper.Pages].(*tview.Pages)

	form := tview.NewForm().
		AddInputField(helper.Name, "", 20, nil, nil).
		AddDropDown(helper.Currency, []string{"PLN", "EUR", "USD", "AUD"}, 0, nil).
		AddInputField(helper.InitialAmount, "", 20, func(textToCheck string, lastChar rune) bool {
			floatValue, err := strconv.ParseFloat(textToCheck, 64)
			if err != nil {
				return false
			}
			if floatValue < 0 {
				return false
			}
			return true
		}, nil).
		AddCheckbox(helper.UseInBudget, true, nil)

	form.SetBorder(true).SetTitle("New Account").SetTitleAlign(tview.AlignLeft)

	form.AddButton("Save", func() {
		//Get form values
		accountName := form.GetFormItemByLabel(helper.Name).(*tview.InputField).GetText()
		_, currency := form.GetFormItemByLabel(helper.Currency).(*tview.DropDown).GetCurrentOption()
		initialAmount, _ := strconv.ParseFloat(form.GetFormItemByLabel(helper.InitialAmount).(*tview.InputField).GetText(), 64)
		useInBudget := form.GetFormItemByLabel(helper.UseInBudget).(*tview.Checkbox).IsChecked()
		if accountName != "" && currency != "" && form.GetFormItemByLabel(helper.InitialAmount).(*tview.InputField).GetText() != "" {
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
			LoadAppElements(ctx)
			pages.HidePage("createAccount")
			pages.ShowPage("main")
		} else {
			ShowPopup("Fill required fields", helper.Alert, ctx)
		}

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
