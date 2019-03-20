package main

type Dollar struct {
	Amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{Amount: 10}
}

func (d *Dollar) Times(n int) {}
