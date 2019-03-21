package main

import (
	"github.com/jdaman/examples/unit/mocking/interface/mylib"
)

var caller Caller

type Caller struct {
	MylibCaller mylib.LibCallerInterface
}

func main(){
	caller.GreaterThan5()
}

func init() {
	caller.MylibCaller = mylib.New()
}

func New(callerInterface mylib.LibCallerInterface) Caller {
	return Caller{callerInterface}
}

func (c *Caller) GreaterThan5() bool {
	return c.MylibCaller.GetN() > 5
}