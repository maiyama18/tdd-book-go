package main

type Money interface {
	Times(n int) Money
	Equals(m Money) bool
}

type Dollar struct {
	amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{amount: amount}
}
func (d *Dollar) Times(n int) *Dollar {
	return NewDollar(d.amount * n)
}
func (d *Dollar) Equals(o *Dollar) bool {
	return d.amount == o.amount
}

type Franc struct {
	amount int
}

func NewFranc(amount int) *Franc {
	return &Franc{amount: amount}
}
func (f *Franc) Times(n int) *Franc {
	return NewFranc(f.amount * n)
}
func (f *Franc) Equals(o *Franc) bool {
	return f.amount == o.amount
}
