package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupAndTeardownTest(t *testing.T) func() {
	t.Log("Setup Test")
	return func() {
		t.Log("Teardown Test")
	}
}

func TestMultiply(t *testing.T) {
	t.Cleanup(setupAndTeardownTest(t))

	t.Log("Testing now Multiply")
	res := multiply(3, 3)
	assert.Equal(t, 9, res)
}
