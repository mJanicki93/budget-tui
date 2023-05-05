package cmd

import (
	"budgettui/pkg/budget"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "client",
	Short: "Run CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.Open("data.json")
		if err != nil {
			err := budget.CreateNewDataFile()
			if err != nil {
				fmt.Println(err)
				return
			}
			data, _ := budget.OpenFile()
			fmt.Println("Please create new budget")
			data.CreateNewBudget()
			return
		}
		byteValue, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		data := budget.Data{}
		err = json.Unmarshal(byteValue, &data)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(data)
	},
}
