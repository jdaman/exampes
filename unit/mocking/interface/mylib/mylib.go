package mylib

type LibCallerInterface interface {
	GetN() int
}

type LibCaller struct {}

func New() LibCaller {
	return LibCaller{}
}

func (LibCaller) GetN() int {
	return 8
}