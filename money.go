package main

type Expression interface {
	Reduce(to string) *Money
}

type Money struct {
	Amount   int
	Currency string
}

func NewMoney(amount int, currency string) *Money {
	return &Money{Amount: amount, Currency: currency}
}
func NewDollar(amount int) *Money {
	return NewMoney(amount, "USD")
}
func NewFranc(amount int) *Money {
	return NewMoney(amount, "CHF")
}

func (m *Money) Plus(o *Money) Expression {
	return NewSum(m, o)
}
func (m *Money) Times(n int) *Money {
	return &Money{Amount: n * m.Amount, Currency: m.Currency}
}
func (m *Money) Reduce(to string) *Money {
	return m
}
func (m *Money) Equals(o *Money) bool {
	return m.Currency == o.Currency && m.Amount == o.Amount
}

type Sum struct {
	Augend *Money
	Addend *Money
}

func NewSum(augend, addend *Money) *Sum {
	return &Sum{Augend: augend, Addend: addend}
}

func (s *Sum) Reduce(to string) *Money {
	amount := s.Augend.Amount + s.Addend.Amount
	return NewMoney(amount, to)
}

type Bank struct{}

func NewBank() *Bank {
	return &Bank{}
}

func (b *Bank) Reduce(src Expression, to string) *Money {
	return src.Reduce(to)
}
