package budget

import "time"

type Expanse struct {
	ID     uint
	Name   string
	Amount float32
	Date   time.Time
}

func (e Expanse) Plan(name string, amount float32, date time.Time) {

}
