package cmd

import (
	"budgetcli/pkg/budget"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
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
		file, err := os.Open("data.json")
		if err != nil {
			err := budget.CreateNewDataFile()
			if err != nil {
				fmt.Println(err)
				return
			}
			file, err = os.Open("data.json")
			if err != nil {
				fmt.Println(err)
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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
