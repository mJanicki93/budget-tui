package budget

type Account struct {
	ID           uint
	Name         string
	Currency     string
	Transactions []Transaction
}
