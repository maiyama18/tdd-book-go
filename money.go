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

type Franc struct {
	amount int
}

func NewFranc(amount int) *Franc {
	return &Franc{amount: amount}
}
func (f *Franc) Times(n int) *Franc {
	return NewFranc(f.amount * n)
}
