package tui

import (
	"budgettui/pkg/budget"
	"budgettui/pkg/helper"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func GetHelpWindow(ctx budget.Context) *tview.Frame {
	pages := ctx[helper.Pages].(*tview.Pages)

	help1 := tview.NewTextView().
		SetDynamicColors(true).
		SetText(`[green]Navigation

[yellow]Left arrow[white]: Move left.
[yellow]Right arrow[white]: Move right.
[yellow]Down arrow[white]: Move down.
[yellow]Up arrow[white]: Move up.
[yellow]Ctrl-A[white]: Create new account.
[yellow]Ctrl-B[white]: Create new budget.
[yellow]Ctrl-T[white]: Make transfer.
[yellow]Ctrl-R[white]: Reload data.

[blue]Press Enter for more help, press Escape to return.`)
	help2 := tview.NewTextView().
		SetDynamicColors(true).
		SetText(`[green]Editing[white]

Type to enter text.
[yellow]Ctrl-H, Backspace[white]: Delete the left character.
[yellow]Ctrl-D, Delete[white]: Delete the right character.
[yellow]Ctrl-K[white]: Delete until the end of the line.
[yellow]Ctrl-W[white]: Delete the rest of the word.
[yellow]Ctrl-U[white]: Delete the current line.

[blue]Press Enter for more help, press Escape to return.`)
	help3 := tview.NewTextView().
		SetDynamicColors(true).
		SetText(`[green]Selecting Text[white]

Move while holding Shift or drag the mouse.
Double-click to select a word.

[green]Clipboard

[yellow]Ctrl-Q[white]: Copy.
[yellow]Ctrl-X[white]: Cut.
[yellow]Ctrl-V[white]: Paste.
		
[green]Undo

[yellow]Ctrl-Z[white]: Undo.
[yellow]Ctrl-Y[white]: Redo.

[blue]Press Enter for more help, press Escape to return.`)

	help := tview.NewFrame(help1).
		SetBorders(1, 1, 0, 0, 2, 2)
	help.SetBorder(true).
		SetTitle("Help").
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyEscape {
				pages.SwitchToPage("main")
				return nil
			} else if event.Key() == tcell.KeyEnter {
				switch {
				case help.GetPrimitive() == help1:
					help.SetPrimitive(help2)
				case help.GetPrimitive() == help2:
					help.SetPrimitive(help3)
				case help.GetPrimitive() == help3:
					help.SetPrimitive(help1)
				}
				return nil
			}
			return event
		})

	//helpPages := []*tview.TextView{help1, help2, help3}

	return help
}
