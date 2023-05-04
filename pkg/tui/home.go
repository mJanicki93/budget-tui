package tui

import "github.com/rivo/tview"

func Home() {
	app := tview.NewApplication()

	helpInfo := tview.NewTextView()
	helpInfo.SetBorder(true)
	helpInfo.
		SetText(" Press F1 for help, press Ctrl-C to exit")

	budgetInfo := tview.NewTextView()
	budgetInfo.SetBorder(true)
	budgetInfo.
		SetText(" Budget info")

	someInfo := tview.NewBox()
	someInfo.SetBorder(true)

	menu := tview.NewList().
		AddItem("Wspólne", "2 000 PLN", '1', func() {
			app.SetFocus(someInfo)
		}).
		AddItem("eMax", "3 000 PLN", '2', nil).
		AddItem("BNP", "5 000 PLN", '3', nil).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})

	mainView := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 0).
		AddItem(helpInfo, 2, 0, 1, 2, 0, 0, false).
		AddItem(menu, 1, 0, 1, 1, 0, 0, true).
		AddItem(someInfo, 1, 1, 1, 1, 0, 0, false).
		AddItem(budgetInfo, 0, 0, 1, 2, 0, 0, false)

	if err := app.SetRoot(mainView,
		true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
