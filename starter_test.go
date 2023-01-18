package starter_test

import (
	"testing"

	starter "github.com/aujenya/go_test_starter"
	"github.com/stretchr/testify/assert"
)

func TestSayHello(t *testing.T) {
	greeting := starter.SayHello("Aurélien")
	assert.Equal(t, "Hello Aurélien. Welcome!", greeting)

	another_greeting := starter.SayHello("Jennifer")
	assert.Equal(t, "Hello Jennifer. Welcome!", another_greeting)
}

func TestOddOrEven(t *testing.T) {
	t.Run("Check Non Negative Numbers", func(t *testing.T) {
		assert.Equal(t, "45 is an odd number", starter.OddOrEven(45))
		assert.Equal(t, "42 is an even number", starter.OddOrEven(42))
		assert.Equal(t, "0 is an even number", starter.OddOrEven(0))
	})
	t.Run("Check Negative Numbers", func(t *testing.T) {
		assert.Equal(t, "-45 is an odd number", starter.OddOrEven(-45))
		assert.Equal(t, "-42 is an even number", starter.OddOrEven(-42))
	})
}
