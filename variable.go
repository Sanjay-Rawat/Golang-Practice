package main

import (
	"fmt"
	"reflect"
)

func VariableDeclairation() {
	var a int = 100
	println(a)
	b := 200
	println(b)
	flag := true
	println(flag)
	c := 3.2
	println(c)
	fmt.Println("type of c is ", reflect.TypeOf(c))

	fmt.Printf("type of c is %T\n", c)
	var s interface{} = "string"
	if _, ok := s.(float64); ok {
		fmt.Println("s is a string")
	} else {
		fmt.Println("s is not a string")
	}
}
