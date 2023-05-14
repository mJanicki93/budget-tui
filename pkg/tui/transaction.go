package tui

import (
	"budgettui/pkg/budget"
	"budgettui/pkg/helper"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strconv"
	"strings"
	"time"
)

func GetNewTransactionForm(accountID uint, pageName string, income bool, ctx budget.Context) *tview.Form {
	data := ctx[helper.Data].(*budget.Data)

	pages := ctx[helper.Pages].(*tview.Pages)

	accountNames := func() []string {
		var accountNamesList []string
		for _, account := range data.Budgets[data.CurrentBudgetID].Accounts {
			accountNamesList = append(accountNamesList, account.Name)
		}
		return accountNamesList
	}

	currentDate := time.Now().Format(time.DateOnly)

	form := tview.NewForm().
		AddInputField(helper.Description, "", 20, nil, nil).
		AddDropDown(helper.Category, data.Budgets[data.CurrentBudgetID].Categories, 0, nil).
		AddInputField(helper.Amount, "", 20, func(textToCheck string, lastChar rune) bool {
			intValue, err := strconv.ParseFloat(textToCheck, 64)
			if err != nil {
				return false
			}
			if intValue < 1 {
				return false
			}
			return true
		}, nil).
		AddInputField("Date", currentDate, 20, func(textToCheck string, lastChar rune) bool {
			converted := strings.Split(textToCheck, "")
			for i, v := range converted {
				if i == 0 || i == 1 || i == 2 || i == 3 || i == 5 || i == 6 || i == 8 || i == 9 {
					_, err := strconv.Atoi(v)
					if err != nil {
						return false
					}
				}
				if i == 4 || i == 7 {
					if v != "-" {
						return false
					}
				}
			}
			if len(textToCheck) == 10 {
				_, err := time.Parse(time.DateOnly, textToCheck)
				if err != nil {
					return false
				}
			}
			if len(textToCheck) > 10 {
				return false
			}
			return true

		}, nil).
		AddDropDown(helper.Account, accountNames(), int(accountID), nil)

	if income {
		form.SetBorder(true).SetTitle("Income").SetTitleAlign(tview.AlignLeft).SetBorderColor(tcell.ColorDarkGreen)
	} else {
		form.SetBorder(true).SetTitle("Outcome").SetTitleAlign(tview.AlignLeft).SetBorderColor(tcell.ColorDarkRed)
	}

	form.AddButton("Save", func() {
		//Get form values
		amount, _ := strconv.ParseFloat(form.GetFormItemByLabel(helper.Amount).(*tview.InputField).GetText(), 64)
		description := form.GetFormItemByLabel(helper.Description).(*tview.InputField).GetText()
		_, category := form.GetFormItemByLabel(helper.Category).(*tview.DropDown).GetCurrentOption()
		i, _ := form.GetFormItemByLabel(helper.Account).(*tview.DropDown).GetCurrentOption()
		date := form.GetFormItemByLabel("Date").(*tview.InputField).GetText()
		timeDate, _ := time.Parse(time.DateOnly, date)

		if description != "" && form.GetFormItemByLabel(helper.Amount).(*tview.InputField).GetText() != "" {
			//Add budget entity
			var transaction budget.Transaction
			if income {
				transaction = budget.Income{
					Description: description,
					Amount:      amount,
					Category:    category,
					Date:        timeDate,
				}
			} else {
				transaction = budget.Expanse{
					Description: description,
					Amount:      amount,
					Category:    category,
					Date:        timeDate,
				}
			}

			budget.CommitTransaction(transaction, uint(i), ctx)

			//Actions
			LoadAppElements(ctx)
			pages.HidePage(pageName)
			pages.ShowPage("main")

		} else {
			ShowPopup("Fill required fields", helper.Alert, ctx)
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

func GetQuickTransactionForm(income bool, ctx budget.Context) *tview.Form {
	pageName := "quickOutcome"
	if income {
		pageName = "quickIncome"
	}
	pages := ctx[helper.Pages].(*tview.Pages)
	form := tview.NewForm().
		AddInputField(helper.Description, "", 20, nil, nil).
		AddInputField(helper.Amount, "", 20, func(textToCheck string, lastChar rune) bool {
			intValue, err := strconv.ParseFloat(textToCheck, 64)
			if err != nil {
				return false
			}
			if intValue < 1 {
				return false
			}
			return true
		}, nil)

	if income {
		form.SetBorder(true).SetTitle("Income").SetTitleAlign(tview.AlignLeft).SetBorderColor(tcell.ColorDarkGreen)
	} else {
		form.SetBorder(true).SetTitle("Outcome").SetTitleAlign(tview.AlignLeft).SetBorderColor(tcell.ColorDarkRed)
	}

	form.AddButton("Save", func() {
		amount, _ := strconv.ParseFloat(form.GetFormItemByLabel(helper.Amount).(*tview.InputField).GetText(), 64)
		description := form.GetFormItemByLabel(helper.Description).(*tview.InputField).GetText()
		if description != "" && form.GetFormItemByLabel(helper.Amount).(*tview.InputField).GetText() != "" {
			//Add budget entity
			var transaction budget.Transaction
			if income {
				transaction = budget.Income{
					Description: description,
					Amount:      amount,
					Category:    "",
					Date:        time.Now(),
				}
			} else {
				transaction = budget.Expanse{
					Description: description,
					Amount:      amount,
					Category:    "",
					Date:        time.Now(),
				}
			}

			budget.CommitTransaction(transaction, 0, ctx)

			//Actions
			LoadAppElements(ctx)
			form.GetFormItemByLabel(helper.Amount).(*tview.InputField).SetText("")
			form.GetFormItemByLabel(helper.Description).(*tview.InputField).SetText("")
			form.SetFocus(0)
			pages.HidePage(pageName)
			pages.ShowPage("main")
		} else {
			ShowPopup("Fill required fields", helper.Alert, ctx)
		}

	})
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			form.GetFormItemByLabel(helper.Amount).(*tview.InputField).SetText("")
			form.GetFormItemByLabel(helper.Description).(*tview.InputField).SetText("")
			form.SetFocus(0)
			pages.HidePage(pageName)
			pages.ShowPage("main")
		}
		return event
	})

	return form
}

func GetTransferForm(ctx budget.Context) *tview.Form {
	data := ctx[helper.Data].(*budget.Data)
	pages := ctx[helper.Pages].(*tview.Pages)

	currentDate := time.Now().Format(time.DateOnly)

	accountNames := func() []string {
		var accountNamesList []string
		for _, account := range data.Budgets[data.CurrentBudgetID].Accounts {
			accountNamesList = append(accountNamesList, account.Name)
		}
		return accountNamesList
	}

	form := tview.NewForm().
		AddInputField(helper.Description, "", 20, nil, nil).
		AddDropDown(helper.Category, data.Budgets[data.CurrentBudgetID].Categories, 0, nil).
		AddInputField(helper.Amount, "", 20, func(textToCheck string, lastChar rune) bool {
			intValue, err := strconv.ParseFloat(textToCheck, 64)
			if err != nil {
				return false
			}
			if intValue < 1 {
				return false
			}
			return true
		}, nil).
		AddInputField("Date", currentDate, 20, func(textToCheck string, lastChar rune) bool {
			converted := strings.Split(textToCheck, "")
			for i, v := range converted {
				if i == 0 || i == 1 || i == 2 || i == 3 || i == 5 || i == 6 || i == 8 || i == 9 {
					_, err := strconv.Atoi(v)
					if err != nil {
						return false
					}
				}
				if i == 4 || i == 7 {
					if v != "-" {
						return false
					}
				}
			}
			if len(textToCheck) == 10 {
				_, err := time.Parse(time.DateOnly, textToCheck)
				if err != nil {
					return false
				}
			}
			if len(textToCheck) > 10 {
				return false
			}
			return true

		}, nil).
		AddDropDown("From", accountNames(), 0, nil).
		AddDropDown("To", accountNames(), 0, nil)

	form.SetBorder(true).SetTitle("Outcome").SetTitleAlign(tview.AlignLeft).SetBorderColor(tcell.ColorDarkRed)
	form.AddButton("Save", func() {
		//Get form values
		amount, _ := strconv.ParseFloat(form.GetFormItemByLabel(helper.Amount).(*tview.InputField).GetText(), 64)
		description := form.GetFormItemByLabel(helper.Description).(*tview.InputField).GetText()
		_, category := form.GetFormItemByLabel(helper.Category).(*tview.DropDown).GetCurrentOption()
		fromID, fromName := form.GetFormItemByLabel("From").(*tview.DropDown).GetCurrentOption()
		toID, toName := form.GetFormItemByLabel("To").(*tview.DropDown).GetCurrentOption()
		date := form.GetFormItemByLabel("Date").(*tview.InputField).GetText()
		timeDate, _ := time.Parse(time.DateOnly, date)

		if description != "" && form.GetFormItemByLabel(helper.Amount).(*tview.InputField).GetText() != "" && fromID != toID {
			expanse := budget.Expanse{
				Description: fmt.Sprintf("%s (%s)", description, toName),
				Amount:      amount,
				Category:    category,
				Date:        timeDate,
			}

			income := budget.Income{
				Description: fmt.Sprintf("%s (%s)", description, fromName),
				Amount:      amount,
				Category:    category,
				Date:        timeDate,
			}

			budget.CommitTransaction(expanse, uint(fromID), ctx)
			budget.CommitTransaction(income, uint(toID), ctx)

			//Actions
			LoadAppElements(ctx)
			pages.HidePage("transferForm")
			pages.ShowPage("main")
		} else if fromID == toID {
			ShowPopup("Same account in both fields", helper.Alert, ctx)
		} else {
			ShowPopup("Fill required fields", helper.Alert, ctx)
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

func GetTransactionForm(accountID uint, transactionID uint, pageName string, ctx budget.Context) *tview.Form {
	data := ctx[helper.Data].(*budget.Data)

	pages := ctx[helper.Pages].(*tview.Pages)
	currentTransaction := data.Budgets[data.CurrentBudgetID].Accounts[accountID].Transactions[transactionID]
	categoryIndex := 0
	for i, category := range data.Budgets[data.CurrentBudgetID].Categories {
		if category == currentTransaction.Category {
			categoryIndex = i
		}

	}
	form := tview.NewForm().
		AddInputField(helper.Description, currentTransaction.Description, 20, nil, nil).
		AddDropDown(helper.Category, data.Budgets[data.CurrentBudgetID].Categories, categoryIndex, nil).
		AddInputField(helper.Amount, fmt.Sprintf("%.2f", currentTransaction.Amount), 20, func(textToCheck string, lastChar rune) bool {
			_, err := strconv.ParseFloat(textToCheck, 64)
			if err != nil {
				return false
			}
			return true
		}, nil).
		AddInputField("Date", currentTransaction.Date.Format(time.DateOnly), 20, func(textToCheck string, lastChar rune) bool {
			converted := strings.Split(textToCheck, "")
			for i, v := range converted {
				if i == 0 || i == 1 || i == 2 || i == 3 || i == 5 || i == 6 || i == 8 || i == 9 {
					_, err := strconv.Atoi(v)
					if err != nil {
						return false
					}
				}
				if i == 4 || i == 7 {
					if v != "-" {
						return false
					}
				}
			}
			if len(textToCheck) == 10 {
				_, err := time.Parse(time.DateOnly, textToCheck)
				if err != nil {
					return false
				}
			}
			if len(textToCheck) > 10 {
				return false
			}
			return true

		}, nil)

	form.SetBorder(true).SetTitle("Outcome").SetTitleAlign(tview.AlignLeft).SetBorderColor(tcell.ColorDarkRed)
	form.AddButton("Save", func() {
		//Get form values
		amount, _ := strconv.ParseFloat(form.GetFormItemByLabel(helper.Amount).(*tview.InputField).GetText(), 64)
		description := form.GetFormItemByLabel(helper.Description).(*tview.InputField).GetText()
		_, category := form.GetFormItemByLabel(helper.Category).(*tview.DropDown).GetCurrentOption()
		date := form.GetFormItemByLabel("Date").(*tview.InputField).GetText()
		timeDate, _ := time.Parse(time.DateOnly, date)

		if description != "" && form.GetFormItemByLabel(helper.Amount).(*tview.InputField).GetText() != "" {
			//Add budget entity
			newTransaction := budget.TransactionEntity{
				ID:          uint(transactionID),
				Description: description,
				Amount:      amount,
				Date:        timeDate,
				Category:    category,
			}
			budget.EditTransaction(uint(accountID), newTransaction, ctx)
			//Actions
			LoadAppElements(ctx)
			pages.HidePage(pageName)
			pages.ShowPage("main")
		} else {
			ShowPopup("Fill required fields", helper.Alert, ctx)
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
