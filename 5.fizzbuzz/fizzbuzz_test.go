package fizzbuzz

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	assert.Equal(t, "fizz", FizzBuzz(3))
	assert.Equal(t, "fizz", FizzBuzz(6))
	assert.Equal(t, "buzz", FizzBuzz(5))
	assert.Equal(t, "buzz", FizzBuzz(10))
	assert.Equal(t, "fizzbuzz", FizzBuzz(15))
	assert.Equal(t, "fizzbuzz", FizzBuzz(30))
}
