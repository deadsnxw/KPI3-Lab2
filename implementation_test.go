package lab2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEasy1(t *testing.T) {
	result := ConvertPostfixToPrefix("2 3 + 5 *")
	assert.Equal(t, "* + 2 3 5", result, "1")
}

func TestEasy2(t *testing.T) {
	result := ConvertPostfixToPrefix("3 8 ^ 6 +")
	assert.Equal(t, "+ 6 ^ 3 8", result, "2")
}

func TestHard1(t *testing.T) {
	result := ConvertPostfixToPrefix("4 7 + 5 * 3 7 - / 5 6 + ^")
	assert.Equal(t, "^ / * + 4 7 5 - 3 7 + 5 6", result, "3")
}

func TestHard2(t *testing.T) {
	result := ConvertPostfixToPrefix("3 5 + 8 - 7 * 4 2 ^ 5 + /")
	assert.Equal(t, "/ * - 8 + 3 5 7 + 5 ^ 4 2", result, "4")
}
