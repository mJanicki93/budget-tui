package budget

import "time"

type Transaction struct {
	ID          uint
	Description string
	Amount      float64
	Category    string
	Date        time.Time
}
