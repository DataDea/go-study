package myintf

import (
	"fmt"
)

type FuncCaller func(interface{}) interface{}

func (f FuncCaller) call(p interface{}) interface{} {
	return f(p)
}

func incr(v interface{}) interface{} {
	return v.(int) + 1
}

func Myintf2() {
	var i FuncCaller = incr
	v := i.call(2)
	fmt.Println(v)

}
