package mylib

type CallerInterface interface {
	GetN() int
}

type Caller struct {}

func New() Caller {
	return Caller{}
}

func (Caller) GetN() int {
	return 8
}