package main

import (
	"testing"
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
	if res != 9 {
		t.Fatal("Res must be 9")
	}
}

func TestTableDoesntWork(t *testing.T) {
	name := "this does not work works"
	type tableTest struct {
		name   *string
		val1   int
		val2   int
		result int
	}
	testDefinition := []tableTest{
		{
			val1:   1,
			val2:   2,
			result: 2,
		},
		{
			name:   &name,
			val1:   1,
			val2:   2,
			result: 2,
		},
	}

	for _, test := range testDefinition {
		// Nill pointer at this point for test1
		t.Run(*test.name, func(t *testing.T) {
			res := multiply(test.val1, test.val2)
			if res != test.result {
				t.Fatal("Res must be", test.result)
			}
		})
	}
}

func TestTable(t *testing.T) {
	name := "this works"
	type tableTest struct {
		name   *string
		val1   int
		val2   int
		result int
	}
	testDefinition := []tableTest{
		{
			name:   &name,
			val1:   1,
			val2:   2,
			result: 2,
		},
		{
			name:   &name,
			val1:   1,
			val2:   2,
			result: 2,
		},
	}

	for _, test := range testDefinition {
		t.Run(*test.name, func(t *testing.T) {
			t.Log("Testing now Multiply", test.name)
			res := multiply(test.val1, test.val2)
			if res != test.result {
				t.Fatal("Res must be", test.result)
			}
		})
	}
}
