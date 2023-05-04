package budget

import "time"

type Income struct {
	ID     uint
	Name   string
	Amount float32
	Date   time.Time
}
