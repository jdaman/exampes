package main

import (
	"testing"
)

// When writing our tests we create a struct matching our interface owning the function to be mocked
type MockCaller struct {}

// The mocked function should simply call a function var that we can access within our tests.
func (MockCaller) GetN() int {
	return stubGetN()
}
var stubGetN func() int

// Within the test our function var definition is altered to make it behave how we want for our testing purpose.
func TestCaller_GreaterThan5(t *testing.T) {
	caller := New(MockCaller{})

	// Test when greater
	stubGetN = func() int {
		return 10
	}

	if !caller.GreaterThan5() {
		t.Errorf("expected true but got false")
	}

	// Test when lesser
	stubGetN = func() int {
		return 4
	}

	if caller.GreaterThan5() {
		t.Errorf("expected false but got true")
	}
}

func TestLessThan5(t *testing.T) {
	// Test when greater
	stubGetN = func() int {
		return 7
	}

	if LessThan5(MockCaller{}) {
		t.Errorf("expected false but got true")
	}

	// Test when lesser
	stubGetN = func() int {
		return 2
	}

	if !LessThan5(MockCaller{}) {
		t.Errorf("expected true but got false")
	}
}