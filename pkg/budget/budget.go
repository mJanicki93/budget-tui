package budget

import "time"

type Budget struct {
	ID         uint
	Name       string
	Settings   Settings
	Accounts   []Account
	Categories []string
	CreatedAt  time.Time
}

func CreateNewBudget() {

}
