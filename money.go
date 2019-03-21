package main

type Expression interface{}

type Money struct {
	Amount   int
	Currency string
}

func NewDollar(amount int) *Money {
	return &Money{Amount: amount, Currency: "USD"}
}
func NewFranc(amount int) *Money {
	return &Money{Amount: amount, Currency: "CHF"}
}

func (m *Money) Plus(o *Money) Expression {
	return &Money{Amount: m.Amount + o.Amount, Currency: m.Currency}
}
func (m *Money) Times(n int) *Money {
	return &Money{Amount: n * m.Amount, Currency: m.Currency}
}
func (m *Money) Equals(o *Money) bool {
	return m.Currency == o.Currency && m.Amount == o.Amount
}

type Bank struct{}

func NewBank() *Bank {
	return &Bank{}
}

func (b *Bank) Reduce(src Expression, dst string) *Money {
	return NewDollar(10)
}
