package main

type Bonus interface {
	isBonus() int
}

type category = string

type Tovar struct {
	Name     string
	Price    int
	Category category
	Bonus    Bonus
}

func (t *Tovar) isBonus() int {
	return t.Bonus.isBonus()
}
