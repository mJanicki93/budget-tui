package cmd

import (
	"budgettui/pkg/tui"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "budget",
	Short: "Budget is simple app for home finances",
	Long: `Simple finance app for:
-managing your home budgets
-planning on every month
-managing debts`,
	Run: func(cmd *cobra.Command, args []string) {
		tui.RunApp()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
