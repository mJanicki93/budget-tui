package tui

import (
	"budgettui/pkg/budget"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strconv"
	"time"
)

func NewCreateBudgetFormInit() *tview.Form {
	form := tview.NewForm().
		AddInputField(Name, "", 20, nil, nil).
		AddDropDown(DefaultCurrency, []string{"PLN", "EUR", "USD", "AUD"}, 0, nil).
		AddInputField(FirstDayOfMonth, "", 2, func(textToCheck string, lastChar rune) bool {
			intValue, err := strconv.Atoi(textToCheck)
			if err != nil {
				return false
			}
			if intValue < 1 || intValue > 31 {
				return false
			}
			return true
		}, nil)

	form.SetBorder(true).SetTitle("New Budget").SetTitleAlign(tview.AlignLeft)

	return form
}

func NewCreateBudgetForm(ctx budget.Context) *tview.Form {
	data, _ := budget.LoadJSONData()
	pages := ctx[Pages].(*tview.Pages)

	form := tview.NewForm().
		AddInputField(Name, "", 20, nil, nil).
		AddDropDown(DefaultCurrency, []string{"PLN", "EUR", "USD", "AUD"}, 0, nil).
		AddInputField(FirstDayOfMonth, "", 2, func(textToCheck string, lastChar rune) bool {
			intValue, err := strconv.Atoi(textToCheck)
			if err != nil {
				return false
			}
			if intValue < 1 || intValue > 31 {
				return false
			}
			return true
		}, nil)

	form.SetBorder(true).SetTitle("New Budget").SetTitleAlign(tview.AlignLeft)
	form.AddButton("Save", func() {
		//Get form values
		firstDayInt, _ := strconv.Atoi(form.GetFormItemByLabel(FirstDayOfMonth).(*tview.InputField).GetText())
		budgetName := form.GetFormItemByLabel(Name).(*tview.InputField).GetText()
		_, defaultCurrency := form.GetFormItemByLabel(DefaultCurrency).(*tview.DropDown).GetCurrentOption()

		if budgetName != "" && form.GetFormItemByLabel(FirstDayOfMonth).(*tview.InputField).GetText() != "" && defaultCurrency == "" {
			//Add budget entity
			data.Budgets = append(data.Budgets, budget.Budget{
				ID:   uint(len(data.Budgets)),
				Name: budgetName,
				Settings: budget.Settings{
					DefaultCurrency: defaultCurrency,
					FirstDay:        uint(firstDayInt),
				},
				CreatedAt: time.Now(),
			})
			data.CurrentBudgetID = uint(len(data.Budgets)) - 1

			//Actions
			_ = data.SaveFile()

			pages.HidePage("createBudget")
			pages.ShowPage("main")
		} else {
			ShowPopup("Fill required fields", Alert, ctx)
		}

	}).AddButton("Quit", func() {
		pages.HidePage("createBudget")
		pages.ShowPage("main")
	})

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			pages.HidePage("createBudget")
			pages.ShowPage("main")
		}
		return event
	})

	return form
}
