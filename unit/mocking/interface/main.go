package main

import (
	"github.com/jdaman/awesomeProject/Mylib"
	"github.com/jdaman/examples/unit/mocking/interface/mylib"
)

// This is an example of how to mock and stub functions through the use of caller interfaces
// In this example we will be mocking the Mylib.GetN function while testing Caller_GreaterThan5 and LessThan5.

var caller Caller

type Caller struct {
	MylibCaller mylib.LibCallerInterface
}

func main(){}

func init() {
	caller.MylibCaller = mylib.New()
}

func New(callerInterface mylib.LibCallerInterface) Caller {
	return Caller{callerInterface}
}

func (c *Caller) GreaterThan5() bool {
	return c.MylibCaller.GetN() > 5
}

func LessThan5(c Mylib.CallerInterface) bool {
	return c.GetN() < 5
}