package main

type Dollar struct {
	amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{amount: amount}
}

func (d *Dollar) Times(n int) *Dollar {
	return NewDollar(d.amount * n)
}
