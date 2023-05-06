package tui

import (
	"budgettui/pkg/budget"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strconv"
	"time"
)

func GetOutcomeForm(data budget.Data, accountID uint, pages *tview.Pages, pageName string, mainMenu *tview.List, accountFrame *tview.Frame, app *tview.Application) *tview.Form {
	accountNames := func() []string {
		var accountNamesList []string
		for _, account := range data.Budgets[data.CurrentBudgetID].Accounts {
			accountNamesList = append(accountNamesList, account.Name)
		}
		return accountNamesList
	}

	form := tview.NewForm().
		AddInputField(Description, "", 20, nil, nil).
		AddDropDown(Category, []string{"", "Home"}, 0, nil).
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

		//Add budget entity
		expanse := budget.Expanse{
			Description: description,
			Amount:      amount,
			Category:    category,
			Date:        time.Now(),
		}

		budget.CommitTransaction(expanse, uint(i))

		//Actions
		LoadMenu(mainMenu, accountFrame, app, pages)
		pages.HidePage(pageName)
		pages.ShowPage("main")
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

func GetIncomeForm(data budget.Data, accountID uint, pages *tview.Pages, pageName string, mainMenu *tview.List, accountFrame *tview.Frame, app *tview.Application) *tview.Form {
	accountNames := func() []string {
		var accountNamesList []string
		for _, account := range data.Budgets[data.CurrentBudgetID].Accounts {
			accountNamesList = append(accountNamesList, account.Name)
		}
		return accountNamesList
	}

	form := tview.NewForm().
		AddInputField(Description, "", 20, nil, nil).
		AddDropDown(Category, []string{"", "Home"}, 0, nil).
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

		//Add budget entity
		income := budget.Income{
			Description: description,
			Amount:      amount,
			Category:    category,
			Date:        time.Now(),
		}

		budget.CommitTransaction(income, uint(i))

		//Actions
		LoadMenu(mainMenu, accountFrame, app, pages)
		pages.HidePage(pageName)
		pages.ShowPage("main")
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

func GetQuickOutcomeForm(pages *tview.Pages, mainMenu *tview.List, accountFrame *tview.Frame, app *tview.Application) *tview.Form {
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

		//Add budget entity
		expanse := budget.Expanse{
			Description: description,
			Amount:      amount,
			Category:    "",
			Date:        time.Now(),
		}

		budget.CommitTransaction(expanse, 0)

		//Actions
		LoadMenu(mainMenu, accountFrame, app, pages)
		form.GetFormItemByLabel(Amount).(*tview.InputField).SetText("")
		form.GetFormItemByLabel(Description).(*tview.InputField).SetText("")
		form.SetFocus(0)
		pages.HidePage("quickOutcome")
		pages.ShowPage("main")
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

func GetQuickIncomeForm(pages *tview.Pages, mainMenu *tview.List, accountFrame *tview.Frame, app *tview.Application) *tview.Form {
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

		//Add budget entity
		income := budget.Income{
			Description: description,
			Amount:      amount,
			Category:    "",
			Date:        time.Now(),
		}

		budget.CommitTransaction(income, 0)

		//Actions
		LoadMenu(mainMenu, accountFrame, app, pages)
		form.GetFormItemByLabel(Amount).(*tview.InputField).SetText("")
		form.GetFormItemByLabel(Description).(*tview.InputField).SetText("")
		form.SetFocus(0)
		pages.HidePage("quickIncome")
		pages.ShowPage("main")
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

func GetTransferForm(data budget.Data, pages *tview.Pages, mainMenu *tview.List, accountFrame *tview.Frame, app *tview.Application) *tview.Form {
	accountNames := func() []string {
		var accountNamesList []string
		for _, account := range data.Budgets[data.CurrentBudgetID].Accounts {
			accountNamesList = append(accountNamesList, account.Name)
		}
		return accountNamesList
	}

	form := tview.NewForm().
		AddInputField(Description, "", 20, nil, nil).
		AddDropDown(Category, []string{"", "Home"}, 0, nil).
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
		fromID, _ := form.GetFormItemByLabel("From").(*tview.DropDown).GetCurrentOption()
		toID, _ := form.GetFormItemByLabel("To").(*tview.DropDown).GetCurrentOption()

		//Add budget entity
		expanse := budget.Expanse{
			Description: description,
			Amount:      amount,
			Category:    category,
			Date:        time.Now(),
		}

		income := budget.Income{
			Description: description,
			Amount:      amount,
			Category:    category,
			Date:        time.Now(),
		}

		budget.CommitTransaction(expanse, uint(fromID))
		budget.CommitTransaction(income, uint(toID))

		//Actions
		LoadMenu(mainMenu, accountFrame, app, pages)
		pages.HidePage("transferForm")
		pages.ShowPage("main")
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
