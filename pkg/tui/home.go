package tui

import (
	"budgettui/pkg/budget"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Home() {
	ctx := budget.NewContext()
	app := tview.NewApplication()

	//Creating all views
	pages := tview.NewPages()
	menu := tview.NewList()
	budgetInfo := tview.NewTextView()
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
		App:              app,
		Pages:            pages,
		Menu:             menu,
		BudegtInfo:       budgetInfo,
		DetailsFrame:     detailsFrame,
		MainGrid:         mainGrid,
		HelpInfo:         helpInfo,
		HelpFrame:        helpFrame,
		NewBudgetForm:    newBudgetForm,
		NewAccountForm:   newAccountForm,
		TransferForm:     transferForm,
		QuickOutcomeForm: quickOutcomeForm,
		QuickIncomeForm:  quickIncomeForm,
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
			ShowPopupQuit(Alert, ctx)
			return nil
		}
		return event
	})

	if err := app.SetRoot(pages,
		true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
