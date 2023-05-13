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
	var currentMonth YearMonth
	for _, account := range accounts {
		for _, t := range account.Transactions {
			y, m, d := t.Date.Date()
			if d >= 10 {
				found := findInDates(y, m, finalDates)
				if !found {
					finalDates = append(finalDates, YearMonth{Year: y, Month: m})
				}

			} else {
				currentYear, lastMonth, _ := time.Now().Add(-(10 * 24 * time.Hour)).Date()
				currentMonth = YearMonth{Year: currentYear, Month: lastMonth}
				for _, v := range finalDates {
					if v.Month == lastMonth && v.Year == currentYear {
						continue
					}
				}
				found := findInDates(currentYear, lastMonth, finalDates)
				if !found {
					finalDates = append(finalDates, currentMonth)
				}
			}
		}
	}
	return finalDates
}

func findInDates(y int, m time.Month, listOfDates []YearMonth) bool {
	for _, v := range listOfDates {
		if v.Month == m && v.Year == y {
			return true
		}
	}
	return false
}
