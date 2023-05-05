package tui

import (
	"budgettui/pkg/budget"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strconv"
	"time"
)

func GetOutcomeForm(data budget.Data, accountID uint, pages *tview.Pages, pageName string) *tview.Form {
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

		expanse.NewTransaction(uint(i))

		//Actions

		pages.HidePage(pageName)
		pages.ShowPage("main")
	}).AddButton("Quit", func() {
		pages.HidePage(pageName)
		pages.ShowPage("main")
	})

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyESC {
			pages.HidePage(pageName)
			pages.ShowPage("main")
		}
		return event
	})

	return form
}

func GetIncomeForm(data budget.Data, accountID uint, pages *tview.Pages, pageName string) *tview.Form {
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

		income.NewTransaction(uint(i))

		//Actions

		pages.HidePage(pageName)
		pages.ShowPage("main")
	}).AddButton("Quit", func() {
		pages.HidePage(pageName)
		pages.ShowPage("main")
	})

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyESC {
			pages.HidePage(pageName)
			pages.ShowPage("main")
		}
		return event
	})

	return form
}
