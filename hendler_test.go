package lab2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutput(t *testing.T) {
	handler := MockComputeHandler{}

	handler.Input = "2 2 + 3 *"

	handler.Output = ""

	result, err := handler.Compute()
	if err != nil {
	}

	assert.Equal(t, "* + 2 2 3", result, "Output test failed")
}

func TestError(t *testing.T) {
	handler := MockComputeHandler{}

	handler.Input = "2 + 2 * 3"

	handler.Output = ""

	result, err := handler.Compute()
	if result != "" {
	}
	assert.True(t, err != nil, "Error test completed")
}
