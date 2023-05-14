package budget

import (
	"budgettui/pkg/helper"
	"time"
)

type TransactionEntity struct {
	ID          uint
	Description string
	Amount      float64
	Category    string
	Date        time.Time
}

type Transaction interface {
	NewTransaction(accountID uint, ctx Context)
	UpdateAccount(accountID uint, ctx Context)
}

func CommitTransaction(t Transaction, id uint, ctx Context) {
	t.NewTransaction(id, ctx)
	t.UpdateAccount(id, ctx)
}

func EditTransaction(accountID uint, newT TransactionEntity, ctx Context) {
	data := ctx[helper.Data].(*Data)
	data.Budgets[data.CurrentBudgetID].Accounts[accountID].Transactions[newT.ID] = newT
	data.Budgets[data.CurrentBudgetID].Accounts[accountID].Transactions[newT.ID].Date = newT.Date
	_ = data.SaveFile()
}
