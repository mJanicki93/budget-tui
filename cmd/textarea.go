package cmd

import (
	"budgettui/pkg/tui"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(textareaCmd)
}

var textareaCmd = &cobra.Command{
	Use:   "textarea",
	Short: "Run CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		tui.TextArea()
	},
}
