package mylib

import "math/rand"

type LibCallerInterface interface {
	GetN() int
}

type LibCaller struct {}

func New() LibCaller {
	return LibCaller{}
}

// Each function to be mocked must be called via an interface.
func (LibCaller) GetN() int {
	return rand.Int()
}