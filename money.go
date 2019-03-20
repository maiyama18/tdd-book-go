package main

type Dollar struct {
	Amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{Amount: amount}
}

func (d *Dollar) Times(n int) *Dollar {
	return NewDollar(d.Amount * n)
}
