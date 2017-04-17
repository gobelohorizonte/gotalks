//
//
//
package main

import (
	"fmt"
)

type Sayer interface {
	Say() string
	ChangeValue() *string
}

type X struct {
	Phrase string
}

func (x *X) Say() string {
	return x.Phrase
}

func (x *X) ChangeValue() *string {
	return &x.Phrase
}

func NewX(s string) *X {
	return &X{s}
}

type Y struct {
	*X
}

func NewY() Y {
	return Y{NewX("Hello World")}
}

func main() {
	y := NewY()
	var sayer Sayer = y
	fmt.Println(y.Say())
	*sayer.ChangeValue() = "Hello Brazil"
	fmt.Println(y.Say())
}
