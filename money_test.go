package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
