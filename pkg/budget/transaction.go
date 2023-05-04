package budget

import "time"

type Transaction struct {
	ID       uint
	Name     string
	Amount   float32
	Category string
	Date     time.Time
}
