package tui

import (
	"budgettui/pkg/budget"
	"budgettui/pkg/helper"
	"github.com/rivo/tview"
	"strconv"
	"strings"
	"time"
)

func Init() {
	app := tview.NewApplication()
	data := budget.Data{}

	//Header
	title := tview.NewTextView().SetDynamicColors(true).SetText(`[green]
	   ___    ________  ___  ___  ________  ________  _______  _________    ___         
	 _|\  \__|\   __  \|\  \|\  \|\   ___ \|\   ____\|\  ___ \|\___   ___\_|\  \__      
	|\   ____\ \  \|\ /\ \  \\\  \ \  \_|\ \ \  \___|\ \   __/\|___ \  \_|\   ____\     
	\ \  \___|\ \   __  \ \  \\\  \ \  \ \\ \ \  \  __\ \  \_|/__  \ \  \\ \  \___|_    
	 \ \_____  \ \  \|\  \ \  \\\  \ \  \_\\ \ \  \|\  \ \  \_|\ \  \ \  \\ \_____  \   
	  \|____|\  \ \_______\ \_______\ \_______\ \_______\ \_______\  \ \__\\|____|\  \  
		____\_\  \|_______|\|_______|\|_______|\|_______|\|_______|   \|__|  ____\_\  \ 
	   |\___    __\                                                         |\___    __\
	   \|___|\__\_|                                                         \|___|\__\_|
			\|__|                                                                \|__|
	`)
	subTitle := tview.NewTextView().SetText("Create first budget").SetTextAlign(1)

	//Get forms
	createBudgetForm := NewCreateBudgetFormInit()
	createAccountForm := NewCreateAccountFormInit()

	//Init pages
	pages := tview.NewPages()

	pages.AddPage("createBudget", tview.NewGrid().
		SetColumns(0, 60, 0).
		SetRows(0, 18, 0).
		AddItem(createBudgetForm, 1, 1, 1, 1, 0, 0, true), true, false)

	pages.AddPage("createAccount", tview.NewGrid().
		SetColumns(0, 60, 0).
		SetRows(0, 14, 0).
		AddItem(createAccountForm, 1, 1, 1, 1, 0, 0, true), true, false)

	//Budget form handler
	createBudgetForm.
		AddButton("Next", func() {
			//Get form values
			firstDayInt, _ := strconv.Atoi(createBudgetForm.GetFormItemByLabel(helper.FirstDayOfMonth).(*tview.InputField).GetText())
			budgetName := createBudgetForm.GetFormItemByLabel(helper.Name).(*tview.InputField).GetText()
			currencyIndex, defaultCurrency := createBudgetForm.GetFormItemByLabel(helper.DefaultCurrency).(*tview.DropDown).GetCurrentOption()
			categories := createBudgetForm.GetFormItemByLabel(helper.Categories).(*tview.TextArea).GetText()

			splitedCategories := strings.Split(categories, ",")
			splitedCategories = append([]string{""}, splitedCategories...)

			if budgetName != "" && createBudgetForm.GetFormItemByLabel(helper.FirstDayOfMonth).(*tview.InputField).GetText() != "" && defaultCurrency != "" && categories != "" {
				//Add budget entity
				data.Budgets = append(data.Budgets, budget.Budget{
					ID:   0,
					Name: budgetName,
					Settings: budget.Settings{
						DefaultCurrency: defaultCurrency,
						FirstDay:        uint(firstDayInt),
					},
					Categories: splitedCategories,
					CreatedAt:  time.Now(),
				})
				data.CurrentBudgetID = 0

				//Actions
				//_ = data.SaveFile()
				createAccountForm.GetFormItemByLabel(helper.Currency).(*tview.DropDown).SetCurrentOption(currencyIndex)
				pages.HidePage("createBudget")
				pages.ShowPage("createAccount")
			} else {
				ctx := budget.Context{helper.Pages: pages}
				ShowPopup("Fill required fields", helper.Alert, ctx)
			}

		})

	//Account form handler
	createAccountForm.
		AddButton("Save", func() {
			//Get form values
			accountName := createAccountForm.GetFormItemByLabel(helper.Name).(*tview.InputField).GetText()
			_, currency := createAccountForm.GetFormItemByLabel(helper.Currency).(*tview.DropDown).GetCurrentOption()
			initialAmount, _ := strconv.ParseFloat(createAccountForm.GetFormItemByLabel(helper.InitialAmount).(*tview.InputField).GetText(), 64)
			useInBudget := createAccountForm.GetFormItemByLabel(helper.UseInBudget).(*tview.Checkbox).IsChecked()
			if accountName != "" && currency != "" && createAccountForm.GetFormItemByLabel(helper.InitialAmount).(*tview.InputField).GetText() != "" {
				//Add account entity
				data.Budgets[data.CurrentBudgetID].Accounts = append(data.Budgets[data.CurrentBudgetID].Accounts, budget.Account{
					ID:          0,
					Name:        accountName,
					Currency:    currency,
					Balance:     initialAmount,
					UseInBudget: useInBudget,
				})

				//Actions
				_ = data.SaveFile()
				app.Stop()
			} else {
				ctx := budget.Context{helper.Pages: pages}
				ShowPopup("Fill required fields", helper.Alert, ctx)
			}

		}).
		AddButton("Quit", func() {
			app.Stop()
		})

	//Main grid
	mainGrid := tview.NewGrid().SetColumns(0, 90, 0).SetRows(12, 5, 0).
		AddItem(title, 0, 1, 1, 1, 0, 0, false).
		AddItem(subTitle, 1, 1, 1, 1, 0, 0, false).
		AddItem(pages, 2, 1, 1, 1, 0, 0, true)

	//Main frame
	mainFrame := tview.NewFrame(mainGrid)

	//Show first page
	pages.ShowPage("createBudget")

	//Run app
	if err := app.SetRoot(mainFrame,
		true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
