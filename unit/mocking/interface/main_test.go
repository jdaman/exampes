package main

import "testing"

type MockCaller struct {}
func (MockCaller) GetN() int {return stubGetN()}
var stubGetN func () int

func TestGreater(t *testing.T) {
	caller := Caller{MylibCaller: MockCaller{}}

	stubGetN = func() int {return 10}

	if !caller.GreaterThan5() {
		t.Errorf("expected true but got false")
	}

}

func TestNotGreater(t *testing.T) {
	caller := Caller{MylibCaller: MockCaller{}}

	stubGetN = func() int {return 4}

	if !caller.GreaterThan5() {
		t.Errorf("expected false but got true")
	}
}