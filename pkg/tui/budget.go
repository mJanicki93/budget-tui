package tui

import (
	"budgettui/pkg/budget"
	"budgettui/pkg/helper"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strconv"
	"strings"
	"time"
)

func NewCreateBudgetFormInit() *tview.Form {
	form := tview.NewForm().
		AddInputField(helper.Name, "", 20, nil, nil).
		AddDropDown(helper.DefaultCurrency, []string{"PLN", "EUR", "USD", "AUD"}, 0, nil).
		AddInputField(helper.FirstDayOfMonth, "", 2, func(textToCheck string, lastChar rune) bool {
			intValue, err := strconv.Atoi(textToCheck)
			if err != nil {
				return false
			}
			if intValue < 1 || intValue > 31 {
				return false
			}
			return true
		}, nil).
		AddTextArea(helper.Categories, "Bills,Savings,Food", 40, 0, 0, nil)

	form.SetBorder(true).SetTitle("New Budget").SetTitleAlign(tview.AlignLeft)

	return form
}

func NewCreateBudgetForm(ctx budget.Context) *tview.Form {
	data := ctx[helper.Data].(*budget.Data)
	pages := ctx[helper.Pages].(*tview.Pages)

	form := tview.NewForm().
		AddInputField(helper.Name, "", 20, nil, nil).
		AddDropDown(helper.DefaultCurrency, []string{"PLN", "EUR", "USD", "AUD"}, 0, nil).
		AddInputField(helper.FirstDayOfMonth, "", 2, func(textToCheck string, lastChar rune) bool {
			intValue, err := strconv.Atoi(textToCheck)
			if err != nil {
				return false
			}
			if intValue < 1 || intValue > 31 {
				return false
			}
			return true
		}, nil).
		AddTextArea(helper.Categories, "Bills,Savings,Food", 40, 0, 0, nil)

	form.SetBorder(true).SetTitle("New Budget").SetTitleAlign(tview.AlignLeft)
	form.AddButton("Save", func() {
		//Get form values
		firstDayInt, _ := strconv.Atoi(form.GetFormItemByLabel(helper.FirstDayOfMonth).(*tview.InputField).GetText())
		budgetName := form.GetFormItemByLabel(helper.Name).(*tview.InputField).GetText()
		_, defaultCurrency := form.GetFormItemByLabel(helper.DefaultCurrency).(*tview.DropDown).GetCurrentOption()
		categories := form.GetFormItemByLabel(helper.Categories).(*tview.TextArea).GetText()

		splitedCategories := strings.Split(categories, ",")
		splitedCategories = append([]string{""}, splitedCategories...)

		if budgetName != "" && form.GetFormItemByLabel(helper.FirstDayOfMonth).(*tview.InputField).GetText() != "" && defaultCurrency == "" && categories != "" {
			//Add budget entity
			data.Budgets = append(data.Budgets, budget.Budget{
				ID:   uint(len(data.Budgets)),
				Name: budgetName,
				Settings: budget.Settings{
					DefaultCurrency: defaultCurrency,
					FirstDay:        uint(firstDayInt),
				},
				Categories: splitedCategories,
				CreatedAt:  time.Now(),
			})
			data.CurrentBudgetID = uint(len(data.Budgets)) - 1

			//Actions
			_ = data.SaveFile()

			pages.HidePage("createBudget")
			pages.ShowPage("main")
		} else {
			ShowPopup("Fill required fields", helper.Alert, ctx)
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
