/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size		   2017
 */

package main

import (
	"fmt"
)

type Sayer interface {
	Say() string
}

type X struct {
	Phrase string
}

func (x X) Say() string {
	return x.Phrase
}

func NewX(s string) X {
	return X{s}
}

type Y struct {
	X
}

func NewY() Y {

	return Y{NewX("Hello World")}
}

/////
///
///
type Cinterface interface {
	ShowC() *string
}

type Conf struct {
	Name string
	Cpf  string
}

func (cc *Conf) ShowC() *string {

	//fmt.Println("opa ShowC1()")
	cc.Name = "Jeff show c just only"
	return &cc.Name
}

func (cc *Conf) ShowC1() *string {

	//fmt.Println("opa ShowC1()")
	cc.Name = "Jeff show c"
	return &cc.Name
}

func (cc *Conf) ShowC2() string {

	//fmt.Println("opa ShowC()")
	cc.Name = "Jeff showc2"
	return cc.Name
}

func (cc *Conf) ShowC3() *Conf {

	cc.Name = "jef showc3 test!!"

	return cc
}

func ShowC4() *Conf {

	//return &Conf{s}
	return &Conf{

		Name: "jeff otoni show 4",
		Cpf:  "393.343.3343-48",
	}
}

func main() {

	var xc Conf

	fmt.Println(*xc.ShowC())
	fmt.Println(*xc.ShowC1())
	fmt.Println(xc.ShowC2())

	fmt.Println(xc.ShowC3().Name)

	fmt.Println("show4: ", ShowC4())
	fmt.Println("show4: ", ShowC4().Name)
	fmt.Println("show4: ", ShowC4().Cpf)

	var sayer Sayer = NewY()
	fmt.Println(sayer.Say())

	sayer = NewX("Hello brazil gophers!")
	fmt.Println(sayer.Say())
}
