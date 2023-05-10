package budget

import (
	"budgettui/pkg/helper"
	"time"
)

type Income struct {
	Name        string
	Description string
	Category    string
	Amount      float64
	Date        time.Time
}

func (e Income) NewTransaction(accountID uint, ctx Context) {
	data := ctx[helper.Data].(*Data)

	newID := 0
	if data.Budgets[data.CurrentBudgetID].Accounts[accountID].Transactions != nil {
		newID = len(data.Budgets[data.CurrentBudgetID].Accounts[accountID].Transactions)
	}
	data.Budgets[data.CurrentBudgetID].Accounts[accountID].Transactions = append(data.Budgets[data.CurrentBudgetID].Accounts[accountID].Transactions, TransactionEntity{
		ID:          uint(newID),
		Description: e.Description,
		Amount:      e.Amount,
		Category:    e.Category,
		Date:        e.Date,
	})
	_ = data.SaveFile()
}

func (e Income) UpdateAccount(accountID uint, ctx Context) {
	data := ctx[helper.Data].(*Data)
	data.Budgets[data.CurrentBudgetID].Accounts[accountID].Balance = data.Budgets[data.CurrentBudgetID].Accounts[accountID].Balance + e.Amount
	_ = data.SaveFile()
}
