package mypkg

import "fmt"

type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParameter(float64)
	MethodWithReturnValue() string
}

type MyType int

func (m MyType) MethodWithoutParameters() {
	fmt.Println("MethodsWithoutParameters Called")
}

func (m MyType) MethodWithParameter(f float64) {
	fmt.Println("MethodWithParameter Called")
}

func (m MyType) MethodWithReturnValue() string {
	fmt.Println("MethodWithReturnValue Called")
	return "Hi from MethodWithReturnValue"
}

func (m MyType) MethodNotInInterface() {
	fmt.Println("MethodNotInInterface Called")
}
