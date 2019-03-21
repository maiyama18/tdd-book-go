package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddition(t *testing.T) {
	five := NewDollar(5)
	sum := five.Plus(five)

	bank := NewBank()
	reduced := bank.Reduce(sum, "USD")

	assert.Equal(t, NewDollar(10), reduced)
}

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	assert.Equal(t, NewDollar(10), five.Times(2))
	assert.Equal(t, NewDollar(15), five.Times(3))
}

func TestEquality(t *testing.T) {
	assert.True(t, NewDollar(5).Equals(NewDollar(5)))
	assert.False(t, NewDollar(5).Equals(NewDollar(10)))
	assert.True(t, NewFranc(5).Equals(NewFranc(5)))
	assert.False(t, NewFranc(5).Equals(NewFranc(10)))
	assert.False(t, NewFranc(5).Equals(NewDollar(5)))
}

func TestCurrency(t *testing.T) {
	assert.Equal(t, "USD", NewDollar(1).Currency)
	assert.Equal(t, "CHF", NewFranc(1).Currency)
}

func TestPlusReturnsSum(t *testing.T) {
	five := NewDollar(5)
	added := five.Plus(five)

	sum, ok := added.(*Sum)
	assert.True(t, ok)
	assert.Equal(t, five, sum.Augend)
	assert.Equal(t, five, sum.Addend)
}

func TestReduceSum(t *testing.T) {
	sum := NewSum(NewDollar(3), NewDollar(4))
	bank := NewBank()

	reduced := bank.Reduce(sum, "USD")
	assert.Equal(t, NewDollar(7), reduced)
}

func TestReduceMoney(t *testing.T) {
	bank := NewBank()
	reduced := bank.Reduce(NewDollar(1), "USD")
	assert.Equal(t, NewDollar(1), reduced)
}

func TestConvertCurrency(t *testing.T) {
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	reduced := bank.Reduce(NewFranc(2), "USD")
	assert.Equal(t, NewDollar(1), reduced)
}
