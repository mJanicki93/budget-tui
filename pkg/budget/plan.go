package budget

import "time"

type Plan struct {
	ID       uint
	Expanses []Expanse
	Incomes  []Income
	DateFrom time.Time
	DateTo   time.Time
}
