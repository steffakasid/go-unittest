package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupTest(t *testing.T) func() {
	log.Println("Setup Test")
	return func() {
		log.Println("Teardown Test")
	}
}

func TestMultiply(t *testing.T) {
	t.Cleanup(setupTest(t))

	log.Println("Testing now Multiply")
	res := multiply(3, 3)
	assert.Equal(t, 9, res)
}
