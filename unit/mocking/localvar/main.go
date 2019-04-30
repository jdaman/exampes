package main

import "github.com/jdaman/examples/unit/mocking/localvar/mylib"

// Define func interface
type GetNFunc func() int

// Instantiate local func var and set to default func definition
var getN GetNFunc = LocalGetN

func main(){}

func LessThan5() bool {
	return getN() < 5
}

// Default func definition
func LocalGetN() int {
	return mylib.GetN()
}

// Create Setter
func SetGetN(nFunc GetNFunc) {
	if nFunc != nil {
		getN = nFunc
	}
}

// Create Getter if needed
func GetGetN() GetNFunc {
	return getN
}