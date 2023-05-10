package budget

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Data struct {
	CurrentBudgetID uint
	Budgets         []Budget
}

type YearMonth struct {
	Year  int
	Month time.Month
}

func CreateNewDataFile() error {
	data := Data{
		CurrentBudgetID: 0,
		Budgets:         nil,
	}

	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile("data.json", file, 0644) //nolint:gosec
	if err != nil {
		return err
	}

	return nil
}

func (d *Data) CreateNewBudget() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Name:")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	name = strings.Replace(name, "\n", "", -1)
	d.Budgets = append(d.Budgets, Budget{
		Name: name,
	})

	err = d.SaveFile()
	if err != nil {
		fmt.Println(err)
	}
}

func (d *Data) SaveFile() error {
	file, err := json.MarshalIndent(d, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile("data.json", file, 0644) //nolint:gosec
	if err != nil {
		return err
	}

	return nil
}

func CheckSaveFile() bool {
	found := true
	_, err := os.Open("data.json")
	if err != nil {
		found = false
	}
	return found
}

func LoadJSONData() *Data {
	data := Data{}
	file, err := os.Open("data.json")
	if err != nil {
		panic(err)
	}
	byteValue, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		panic(err)
	}
	return &data
}

func (d *Data) GetMonths() []YearMonth {
	accounts := d.Budgets[d.CurrentBudgetID].Accounts

	var finalDates []YearMonth

	for _, account := range accounts {
		for _, t := range account.Transactions {
			y, m, d := t.Date.Date()
			if d > 10 {
				finalDates = append(finalDates, YearMonth{Year: y, Month: m})
			}
		}
	}
	return finalDates
}
