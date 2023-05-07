package budget

import "time"

type Transaction struct {
	ID          uint
	Description string
	Amount      float64
	Category    string
	Date        time.Time
}

type transactions interface {
	NewTransaction(accountID uint)
	UpdateAccount(accountID uint)
}

func CommitTransaction(t transactions, id uint) {
	t.NewTransaction(id)
	t.UpdateAccount(id)
}

func EditTransaction(accountID uint, newT Transaction) {
	data, _ := LoadJSONData()
	oldDate := data.Budgets[data.CurrentBudgetID].Accounts[accountID].Transactions[newT.ID].Date
	data.Budgets[data.CurrentBudgetID].Accounts[accountID].Transactions[newT.ID] = newT
	data.Budgets[data.CurrentBudgetID].Accounts[accountID].Transactions[newT.ID].Date = oldDate
	_ = data.SaveFile()
}
