package budget

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Data struct {
	CurrentBudgetID uint
	Budgets         []Budget
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

	err = ioutil.WriteFile("data.json", file, 0644) //nolint:gosec
	if err != nil {
		return err
	}

	return nil
}

func (d Data) CreateNewBudget() {
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

func (d Data) SaveFile() error {
	file, err := json.MarshalIndent(d, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("data.json", file, 0644) //nolint:gosec
	if err != nil {
		return err
	}

	return nil
}
