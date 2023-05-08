package budget

import "time"

type TransactionEntity struct {
	ID          uint
	Description string
	Amount      float64
	Category    string
	Date        time.Time
}

type Transaction interface {
	NewTransaction(accountID uint)
	UpdateAccount(accountID uint)
}

func CommitTransaction(t Transaction, id uint) {
	t.NewTransaction(id)
	t.UpdateAccount(id)
}

func EditTransaction(accountID uint, newT TransactionEntity) {
	data, _ := LoadJSONData()
	oldDate := data.Budgets[data.CurrentBudgetID].Accounts[accountID].Transactions[newT.ID].Date
	data.Budgets[data.CurrentBudgetID].Accounts[accountID].Transactions[newT.ID] = newT
	data.Budgets[data.CurrentBudgetID].Accounts[accountID].Transactions[newT.ID].Date = oldDate
	_ = data.SaveFile()
}
