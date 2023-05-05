package budget

import "time"

type Expanse struct {
	Description string
	Amount      float64
	Category    string
	Date        time.Time
}

func (e Expanse) NewTransaction(accountID uint) {
	data, _ := OpenFile()

	newID := 0
	if data.Budgets[data.CurrentBudgetID].Accounts[accountID].Transactions != nil {
		newID = len(data.Budgets[data.CurrentBudgetID].Accounts[accountID].Transactions)
	}
	data.Budgets[data.CurrentBudgetID].Accounts[accountID].Transactions = append(data.Budgets[data.CurrentBudgetID].Accounts[accountID].Transactions, Transaction{
		ID:          uint(newID),
		Description: e.Description,
		Amount:      -e.Amount,
		Category:    e.Category,
		Date:        e.Date,
	})
	_ = data.SaveFile()

}

func (e Expanse) UpdateAccount(accountID uint) {
	data, _ := OpenFile()
	data.Budgets[data.CurrentBudgetID].Accounts[accountID].Balance = data.Budgets[data.CurrentBudgetID].Accounts[accountID].Balance - e.Amount
	_ = data.SaveFile()
}

func (e Expanse) Plan(name string, amount float32, date time.Time) {

}
