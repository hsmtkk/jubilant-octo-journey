package nthprime_test

import (
	"testing"

	"github.com/hsmtkk/jubilant-octo-journey/nthprime"
	"github.com/stretchr/testify/assert"
)

func TestNthPrime(t *testing.T) {
	assert.Equal(t, 13, nthprime.NthPrime(5))
}

func TestIsPrime(t *testing.T) {
	assert.True(t, nthprime.IsPrime(11))
	assert.False(t, nthprime.IsPrime(12))
}
