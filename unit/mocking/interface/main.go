package main

import (
	"github.com/jdaman/exampes/unit/mocking/interface/mylib"
)

var caller Caller

type Caller struct {
	MylibCaller mylib.CallerInterface
}

func main(){}

func init() {
	caller.MylibCaller = mylib.New()
}

func (Caller) GreaterThan5() bool {
	return caller.MylibCaller.GetN() > 5
}