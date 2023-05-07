package tui

import (
	"budgettui/pkg/budget"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func ShowPopup(text string, style string, ctx budget.Context) {
	pages := ctx[Pages].(*tview.Pages)

	modal := tview.NewModal().SetText(text).AddButtons([]string{"OK"})
	pages.AddPage("modal", modal, true, false)
	modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		pages.RemovePage("modal")
	})
	if style == Alert {
		modal.SetBackgroundColor(tcell.ColorDarkRed)
	}
	if style == Info {
		modal.SetBackgroundColor(tcell.ColorDarkBlue)
	}
	if style == Success {
		modal.SetBackgroundColor(tcell.ColorDarkGreen)
	}
	pages.ShowPage("modal")
}

func ShowPopupYesNo(text string, style string, yesFunction func(), noFunction func(), ctx budget.Context) {
	pages := ctx[Pages].(*tview.Pages)

	modal := tview.NewModal().SetText(text).AddButtons([]string{"YES", "NO"})
	pages.AddPage("modal", modal, true, false)
	modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonLabel == "YES" {
			yesFunction()
		}
		if buttonLabel == "NO" {
			noFunction()
		}
		pages.RemovePage("modal")
	})
	if style == Alert {
		modal.SetBackgroundColor(tcell.ColorDarkRed)
	}
	if style == Info {
		modal.SetBackgroundColor(tcell.ColorDarkBlue)
	}
	if style == Success {
		modal.SetBackgroundColor(tcell.ColorDarkGreen)
	}
	pages.ShowPage("modal")
}

func ShowPopupQuit(style string, ctx budget.Context) {
	pages := ctx[Pages].(*tview.Pages)
	app := ctx[App].(*tview.Application)

	modal := tview.NewModal().SetText("Do you want to quit the application? ").AddButtons([]string{"YES", "NO"})
	pages.AddPage("modal", modal, true, false)
	modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonLabel == "YES" {
			app.Stop()
		}

		pages.RemovePage("modal")
	})
	if style == Alert {
		modal.SetBackgroundColor(tcell.ColorDarkRed)
	}
	if style == Info {
		modal.SetBackgroundColor(tcell.ColorDarkBlue)
	}
	if style == Success {
		modal.SetBackgroundColor(tcell.ColorDarkGreen)
	}
	pages.ShowPage("modal")
}
