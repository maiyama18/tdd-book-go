package main

type Money interface {
	Amount() int
	Times(n int) Money
	Equals(m Money) bool
}

type Dollar struct {
	amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{amount: amount}
}
func (d *Dollar) Amount() int {
	return d.amount
}
func (d *Dollar) Times(n int) Money {
	return NewDollar(d.amount * n)
}
func (d *Dollar) Equals(o Money) bool {
	if _, ok := o.(*Dollar); !ok {
		return false
	}
	return d.amount == o.Amount()
}

type Franc struct {
	amount int
}

func NewFranc(amount int) *Franc {
	return &Franc{amount: amount}
}
func (f *Franc) Amount() int {
	return f.amount
}
func (f *Franc) Times(n int) Money {
	return NewFranc(f.amount * n)
}
func (f *Franc) Equals(o Money) bool {
	if _, ok := o.(*Franc); !ok {
		return false
	}
	return f.amount == o.Amount()
}
