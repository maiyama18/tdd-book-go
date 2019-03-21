package main

type Expression interface {
	Reduce(bank *Bank, to string) *Money
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
func (m *Money) Reduce(bank *Bank, to string) *Money {
	rate := bank.Rate(m.Currency, to)
	return NewMoney(m.Amount/rate, to)
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

func (s *Sum) Reduce(bank *Bank, to string) *Money {
	augendAmount := s.Augend.Amount / bank.Rate(s.Augend.Currency, to)
	addendAmount := s.Addend.Amount / bank.Rate(s.Addend.Currency, to)
	return NewMoney(augendAmount+addendAmount, to)
}

type Bank struct {
	rates map[Pair]int
}

func NewBank() *Bank {
	return &Bank{
		rates: make(map[Pair]int),
	}
}

func (b *Bank) AddRate(from, to string, rate int) {
	pair := NewPair(from, to)
	b.rates[pair] = rate
}
func (b *Bank) Reduce(src Expression, to string) *Money {
	return src.Reduce(b, to)
}
func (b *Bank) Rate(from, to string) int {
	pair := NewPair(from, to)
	rate, ok := b.rates[pair]
	if !ok {
		return 1
	}
	return rate
}

type Pair struct {
	From string
	To   string
}

func NewPair(from, to string) Pair {
	return Pair{From: from, To: to}
}
