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
