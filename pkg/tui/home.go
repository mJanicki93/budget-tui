package tui

import (
	"budgettui/pkg/budget"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Home(data budget.Data) {
	app := tview.NewApplication()

	//Creating all views
	pages := tview.NewPages()
	menu := tview.NewList()
	budgetInfo := tview.NewTextView()
	account := tview.NewFrame(tview.NewGrid())
	helpInfo := tview.NewTextView()
	help := GetHelpWindow(pages)

	newBudget := NewCreateBudgetForm(data, pages)

	newAccount := NewCreateAccountForm(data, pages, menu, account, app)

	transferForm := GetTransferForm(data, pages)

	helpInfo.SetBorder(true)
	helpInfo.
		SetText(" Press F1 for help, press Ctrl-C to exit")

	budgetInfo.SetBorder(true)
	budgetInfo.
		SetText(data.Budgets[data.CurrentBudgetID].Name)

	account.SetBorder(true)
	LoadMenu(menu, account, app, pages)

	mainView := GetMainView(helpInfo, menu, account, budgetInfo)

	pages.AddAndSwitchToPage("main", mainView, true).
		AddPage("help", tview.NewGrid().
			SetColumns(0, 64, 0).
			SetRows(0, 22, 0).
			AddItem(help, 1, 1, 1, 1, 0, 0, true), true, false).
		AddPage("createBudget", tview.NewGrid().
			SetColumns(0, 58, 0).
			SetRows(0, 12, 0).
			AddItem(newBudget, 1, 1, 1, 1, 0, 0, true), true, false).
		AddPage("createAccount", tview.NewGrid().
			SetColumns(0, 58, 0).
			SetRows(0, 13, 0).
			AddItem(newAccount, 1, 1, 1, 1, 0, 0, true), true, false).
		AddPage("transferForm", tview.NewGrid().
			SetColumns(0, 58, 0).
			SetRows(0, 13, 0).
			AddItem(transferForm, 1, 1, 1, 1, 0, 0, true), true, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlR {
			LoadMenu(menu, account, app, pages)
			pages.ShowPage("main")
		}
		if event.Key() == tcell.KeyF1 {
			pages.ShowPage("help")
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
		return event
	})

	if err := app.SetRoot(pages,
		true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
