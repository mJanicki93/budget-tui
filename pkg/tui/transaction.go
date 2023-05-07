package tui

import (
	"budgettui/pkg/budget"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strconv"
	"time"
)

func GetOutcomeForm(accountID uint, pageName string, ctx budget.Context) *tview.Form {
	data, _ := budget.LoadJSONData()

	pages := ctx[Pages].(*tview.Pages)

	accountNames := func() []string {
		var accountNamesList []string
		for _, account := range data.Budgets[data.CurrentBudgetID].Accounts {
			accountNamesList = append(accountNamesList, account.Name)
		}
		return accountNamesList
	}

	form := tview.NewForm().
		AddInputField(Description, "", 20, nil, nil).
		AddDropDown(Category, data.Budgets[data.CurrentBudgetID].Categories, 0, nil).
		AddInputField(Amount, "", 20, func(textToCheck string, lastChar rune) bool {
			intValue, err := strconv.ParseFloat(textToCheck, 64)
			if err != nil {
				return false
			}
			if intValue < 1 {
				return false
			}
			return true
		}, nil).
		AddDropDown(Account, accountNames(), int(accountID), nil)

	form.SetBorder(true).SetTitle("Outcome").SetTitleAlign(tview.AlignLeft).SetBorderColor(tcell.ColorDarkRed)
	form.AddButton("Save", func() {
		//Get form values
		amount, _ := strconv.ParseFloat(form.GetFormItemByLabel(Amount).(*tview.InputField).GetText(), 64)
		description := form.GetFormItemByLabel(Description).(*tview.InputField).GetText()
		_, category := form.GetFormItemByLabel(Category).(*tview.DropDown).GetCurrentOption()
		i, _ := form.GetFormItemByLabel(Account).(*tview.DropDown).GetCurrentOption()

		if description != "" && form.GetFormItemByLabel(Amount).(*tview.InputField).GetText() != "" {
			//Add budget entity
			expanse := budget.Expanse{
				Description: description,
				Amount:      amount,
				Category:    category,
				Date:        time.Now(),
			}
			budget.CommitTransaction(expanse, uint(i))
			//Actions
			LoadAppData(ctx)
			pages.HidePage(pageName)
			pages.ShowPage("main")
		} else {
			ShowPopup("Fill required fields", Alert, ctx)
		}

	}).AddButton("Quit", func() {
		pages.HidePage(pageName)
		pages.ShowPage("main")
	})

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			pages.HidePage(pageName)
			pages.ShowPage("main")
		}
		return event
	})

	return form
}

func GetIncomeForm(accountID uint, pageName string, ctx budget.Context) *tview.Form {
	data, _ := budget.LoadJSONData()

	pages := ctx[Pages].(*tview.Pages)

	accountNames := func() []string {
		var accountNamesList []string
		for _, account := range data.Budgets[data.CurrentBudgetID].Accounts {
			accountNamesList = append(accountNamesList, account.Name)
		}
		return accountNamesList
	}

	form := tview.NewForm().
		AddInputField(Description, "", 20, nil, nil).
		AddDropDown(Category, data.Budgets[data.CurrentBudgetID].Categories, 0, nil).
		AddInputField(Amount, "", 20, func(textToCheck string, lastChar rune) bool {
			intValue, err := strconv.ParseFloat(textToCheck, 64)
			if err != nil {
				return false
			}
			if intValue < 1 {
				return false
			}
			return true
		}, nil).
		AddDropDown(Account, accountNames(), int(accountID), nil)

	form.SetBorder(true).SetTitle("Income").SetTitleAlign(tview.AlignLeft).SetBorderColor(tcell.ColorDarkGreen)
	form.AddButton("Save", func() {
		//Get form values
		amount, _ := strconv.ParseFloat(form.GetFormItemByLabel(Amount).(*tview.InputField).GetText(), 64)
		description := form.GetFormItemByLabel(Description).(*tview.InputField).GetText()
		_, category := form.GetFormItemByLabel(Category).(*tview.DropDown).GetCurrentOption()
		i, _ := form.GetFormItemByLabel(Account).(*tview.DropDown).GetCurrentOption()

		if description != "" && form.GetFormItemByLabel(Amount).(*tview.InputField).GetText() != "" {
			//Add budget entity
			income := budget.Income{
				Description: description,
				Amount:      amount,
				Category:    category,
				Date:        time.Now(),
			}

			budget.CommitTransaction(income, uint(i))

			//Actions
			LoadAppData(ctx)
			pages.HidePage(pageName)
			pages.ShowPage("main")
		} else {
			ShowPopup("Fill required fields", Alert, ctx)
		}

	}).AddButton("Quit", func() {
		pages.HidePage(pageName)
		pages.ShowPage("main")
	})

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			pages.HidePage(pageName)
			pages.ShowPage("main")
		}
		return event
	})

	return form
}

func GetQuickOutcomeForm(ctx budget.Context) *tview.Form {
	pages := ctx[Pages].(*tview.Pages)
	form := tview.NewForm().
		AddInputField(Description, "", 20, nil, nil).
		AddInputField(Amount, "", 20, func(textToCheck string, lastChar rune) bool {
			intValue, err := strconv.ParseFloat(textToCheck, 64)
			if err != nil {
				return false
			}
			if intValue < 1 {
				return false
			}
			return true
		}, nil)

	form.SetBorder(true).SetTitle("Outcome").SetTitleAlign(tview.AlignLeft).SetBorderColor(tcell.ColorDarkRed)
	form.AddButton("Save", func() {
		amount, _ := strconv.ParseFloat(form.GetFormItemByLabel(Amount).(*tview.InputField).GetText(), 64)
		description := form.GetFormItemByLabel(Description).(*tview.InputField).GetText()
		if description != "" && form.GetFormItemByLabel(Amount).(*tview.InputField).GetText() != "" {
			//Add budget entity
			expanse := budget.Expanse{
				Description: description,
				Amount:      amount,
				Category:    "",
				Date:        time.Now(),
			}

			budget.CommitTransaction(expanse, 0)

			//Actions
			LoadAppData(ctx)
			form.GetFormItemByLabel(Amount).(*tview.InputField).SetText("")
			form.GetFormItemByLabel(Description).(*tview.InputField).SetText("")
			form.SetFocus(0)
			pages.HidePage("quickOutcome")
			pages.ShowPage("main")
		} else {
			ShowPopup("Fill required fields", Alert, ctx)
		}

	})
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			form.GetFormItemByLabel(Amount).(*tview.InputField).SetText("")
			form.GetFormItemByLabel(Description).(*tview.InputField).SetText("")
			form.SetFocus(0)
			pages.HidePage("quickOutcome")
			pages.ShowPage("main")
		}
		return event
	})

	return form
}

func GetQuickIncomeForm(ctx budget.Context) *tview.Form {
	pages := ctx[Pages].(*tview.Pages)
	form := tview.NewForm().
		AddInputField(Description, "", 20, nil, nil).
		AddInputField(Amount, "", 20, func(textToCheck string, lastChar rune) bool {
			intValue, err := strconv.ParseFloat(textToCheck, 64)
			if err != nil {
				return false
			}
			if intValue < 1 {
				return false
			}
			return true
		}, nil)

	form.SetBorder(true).SetTitle("Income").SetTitleAlign(tview.AlignLeft).SetBorderColor(tcell.ColorDarkGreen)
	form.AddButton("Save", func() {
		//Get form values
		amount, _ := strconv.ParseFloat(form.GetFormItemByLabel(Amount).(*tview.InputField).GetText(), 64)
		description := form.GetFormItemByLabel(Description).(*tview.InputField).GetText()

		if description != "" && form.GetFormItemByLabel(Amount).(*tview.InputField).GetText() != "" {
			//Add budget entity
			income := budget.Income{
				Description: description,
				Amount:      amount,
				Category:    "",
				Date:        time.Now(),
			}

			budget.CommitTransaction(income, 0)

			//Actions
			LoadAppData(ctx)
			form.GetFormItemByLabel(Amount).(*tview.InputField).SetText("")
			form.GetFormItemByLabel(Description).(*tview.InputField).SetText("")
			form.SetFocus(0)
			pages.HidePage("quickIncome")
			pages.ShowPage("main")
		} else {
			ShowPopup("Fill required fields", Alert, ctx)
		}

	})

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			form.GetFormItemByLabel(Amount).(*tview.InputField).SetText("")
			form.GetFormItemByLabel(Description).(*tview.InputField).SetText("")
			form.SetFocus(0)
			pages.HidePage("quickIncome")
			pages.ShowPage("main")
		}
		return event
	})

	return form
}

func GetTransferForm(ctx budget.Context) *tview.Form {
	data, _ := budget.LoadJSONData()
	pages := ctx[Pages].(*tview.Pages)

	accountNames := func() []string {
		var accountNamesList []string
		for _, account := range data.Budgets[data.CurrentBudgetID].Accounts {
			accountNamesList = append(accountNamesList, account.Name)
		}
		return accountNamesList
	}

	form := tview.NewForm().
		AddInputField(Description, "", 20, nil, nil).
		AddDropDown(Category, data.Budgets[data.CurrentBudgetID].Categories, 0, nil).
		AddInputField(Amount, "", 20, func(textToCheck string, lastChar rune) bool {
			intValue, err := strconv.ParseFloat(textToCheck, 64)
			if err != nil {
				return false
			}
			if intValue < 1 {
				return false
			}
			return true
		}, nil).
		AddDropDown("From", accountNames(), 0, nil).
		AddDropDown("To", accountNames(), 0, nil)

	form.SetBorder(true).SetTitle("Outcome").SetTitleAlign(tview.AlignLeft).SetBorderColor(tcell.ColorDarkRed)
	form.AddButton("Save", func() {
		//Get form values
		amount, _ := strconv.ParseFloat(form.GetFormItemByLabel(Amount).(*tview.InputField).GetText(), 64)
		description := form.GetFormItemByLabel(Description).(*tview.InputField).GetText()
		_, category := form.GetFormItemByLabel(Category).(*tview.DropDown).GetCurrentOption()
		fromID, fromName := form.GetFormItemByLabel("From").(*tview.DropDown).GetCurrentOption()
		toID, toName := form.GetFormItemByLabel("To").(*tview.DropDown).GetCurrentOption()

		if description != "" && form.GetFormItemByLabel(Amount).(*tview.InputField).GetText() != "" && fromID != toID {
			expanse := budget.Expanse{
				Description: fmt.Sprintf("(To:%s)%s", toName, description),
				Amount:      amount,
				Category:    category,
				Date:        time.Now(),
			}

			income := budget.Income{
				Description: fmt.Sprintf("(From:%s)%s", fromName, description),
				Amount:      amount,
				Category:    category,
				Date:        time.Now(),
			}

			budget.CommitTransaction(expanse, uint(fromID))
			budget.CommitTransaction(income, uint(toID))

			//Actions
			LoadAppData(ctx)
			pages.HidePage("transferForm")
			pages.ShowPage("main")
		} else if fromID == toID {
			ShowPopup("Same account in both fields", Alert, ctx)
		} else {
			ShowPopup("Fill required fields", Alert, ctx)
		}
		//Add budget entity

	}).AddButton("Quit", func() {
		pages.HidePage("transferForm")
		pages.ShowPage("main")
	})

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			pages.HidePage("transferForm")
			pages.ShowPage("main")
		}
		return event
	})

	return form
}

func GetTransactionForm(accountID int, transactionID int, pageName string, ctx budget.Context) *tview.Form {
	data, _ := budget.LoadJSONData()

	pages := ctx[Pages].(*tview.Pages)
	currentTransaction := data.Budgets[data.CurrentBudgetID].Accounts[accountID].Transactions[transactionID]
	categoryIndex := 0
	for i, category := range data.Budgets[data.CurrentBudgetID].Categories {
		if category == currentTransaction.Category {
			categoryIndex = i
		}

	}
	form := tview.NewForm().
		AddInputField(Description, currentTransaction.Description, 20, nil, nil).
		AddDropDown(Category, data.Budgets[data.CurrentBudgetID].Categories, categoryIndex, nil).
		AddInputField(Amount, fmt.Sprintf("%.2f", currentTransaction.Amount), 20, func(textToCheck string, lastChar rune) bool {
			_, err := strconv.ParseFloat(textToCheck, 64)
			if err != nil {
				return false
			}
			return true
		}, nil)

	form.SetBorder(true).SetTitle("Outcome").SetTitleAlign(tview.AlignLeft).SetBorderColor(tcell.ColorDarkRed)
	form.AddButton("Save", func() {
		//Get form values
		amount, _ := strconv.ParseFloat(form.GetFormItemByLabel(Amount).(*tview.InputField).GetText(), 64)
		description := form.GetFormItemByLabel(Description).(*tview.InputField).GetText()
		_, category := form.GetFormItemByLabel(Category).(*tview.DropDown).GetCurrentOption()

		if description != "" && form.GetFormItemByLabel(Amount).(*tview.InputField).GetText() != "" {
			//Add budget entity
			newTransaction := budget.Transaction{
				ID:          uint(transactionID),
				Description: description,
				Amount:      amount,
				Category:    category,
			}
			budget.EditTransaction(accountID, newTransaction)
			//Actions
			LoadAppData(ctx)
			pages.HidePage(pageName)
			pages.ShowPage("main")
		} else {
			ShowPopup("Fill required fields", Alert, ctx)
		}

	}).AddButton("Quit", func() {
		pages.HidePage(pageName)
		pages.ShowPage("main")
	})

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			pages.HidePage(pageName)
			pages.ShowPage("main")
		}
		return event
	})

	return form
}
