package budget

type Account struct {
	ID           uint
	Name         string
	Currency     string
	Balance      float64
	UseInBudget  bool
	Transactions []Transaction
}
