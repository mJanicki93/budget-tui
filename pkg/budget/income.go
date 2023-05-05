package budget

import "time"

type Income struct {
	Name        string
	Description string
	Category    string
	Amount      float64
	Date        time.Time
}

func (e Income) NewTransaction(accountID uint) {
	data, _ := OpenFile()

	newID := 0
	if data.Budgets[data.CurrentBudgetID].Accounts[accountID].Transactions != nil {
		newID = len(data.Budgets[data.CurrentBudgetID].Accounts[accountID].Transactions)
	}
	data.Budgets[data.CurrentBudgetID].Accounts[accountID].Transactions = append(data.Budgets[data.CurrentBudgetID].Accounts[accountID].Transactions, Transaction{
		ID:          uint(newID),
		Description: e.Description,
		Amount:      e.Amount,
		Category:    e.Category,
		Date:        e.Date,
	})
	_ = data.SaveFile()

}
