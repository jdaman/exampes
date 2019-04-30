package main

import (
	"testing"
)

// Use Setter in test to mock behavior
func TestLessThan5(t *testing.T) {
	SetGetN(MockGetN0)
	if !LessThan5() {
		t.Errorf("expected false but got true")
	}

	SetGetN(MockGetN9)
	if LessThan5() {
		t.Errorf("expected true but got false")
	}
}

// Mocked function
func MockGetN0() int {
	return 0
}

// Mocked function
func MockGetN9() int {
	return 9
}