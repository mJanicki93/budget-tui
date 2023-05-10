package tui

import (
	"budgettui/pkg/budget"
	"budgettui/pkg/helper"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"time"
)

func Home() {
	ctx := budget.NewContext()
	app := tview.NewApplication()
	data := budget.LoadJSONData()
	y, m, d := time.Now().Date()

	var currentMonth *budget.YearMonth
	if d > int(data.Budgets[data.CurrentBudgetID].Settings.FirstDay) {
		currentMonth = &budget.YearMonth{Year: y, Month: m}
	} else {
		currentYear, lastMonth, _ := time.Now().Add(-(10 * 24 * time.Hour)).Date()
		currentMonth = &budget.YearMonth{Year: currentYear, Month: lastMonth}
	}

	//Creating all views
	pages := tview.NewPages()
	menu := tview.NewList()
	topBar := tview.NewGrid()
	detailsFrame := tview.NewFrame(tview.NewGrid())
	mainGrid := tview.NewGrid()
	helpInfo := tview.NewTextView()
	helpFrame := tview.NewFrame(tview.NewGrid())
	newBudgetForm := tview.NewForm()
	newAccountForm := tview.NewForm()
	transferForm := tview.NewForm()
	quickOutcomeForm := tview.NewForm()
	quickIncomeForm := tview.NewForm()

	//Add to context
	ctx.AddMap(map[string]any{
		helper.App:              app,
		helper.Data:             data,
		helper.CurrentMonth:     currentMonth,
		helper.Pages:            pages,
		helper.Menu:             menu,
		helper.TopBar:           topBar,
		helper.DetailsFrame:     detailsFrame,
		helper.MainGrid:         mainGrid,
		helper.HelpInfo:         helpInfo,
		helper.HelpFrame:        helpFrame,
		helper.NewBudgetForm:    newBudgetForm,
		helper.NewAccountForm:   newAccountForm,
		helper.TransferForm:     transferForm,
		helper.QuickOutcomeForm: quickOutcomeForm,
		helper.QuickIncomeForm:  quickIncomeForm,
	})

	LoadAppElements(ctx)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlR {
			LoadAppElements(ctx)
			pages.ShowPage("main")
		}
		if event.Key() == tcell.KeyF1 {
			pages.ShowPage("helpFrame")
			return nil
		}
		if event.Key() == tcell.KeyCtrlB {
			pages.ShowPage("createBudget")
			return nil
		}
		if event.Key() == tcell.KeyCtrlA {
			pages.ShowPage("createAccount")
			return nil
		}
		if event.Key() == tcell.KeyCtrlT {
			pages.ShowPage("transferForm")
			return nil
		}
		if event.Key() == tcell.KeyCtrlJ {
			pages.ShowPage("quickIncome")
			return nil
		}
		if event.Key() == tcell.KeyCtrlO {
			pages.ShowPage("quickOutcome")
			return nil
		}
		if event.Key() == tcell.KeyCtrlC {
			ShowPopupQuit(helper.Alert, ctx)
			return nil
		}
		return event
	})

	if err := app.SetRoot(pages,
		true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
