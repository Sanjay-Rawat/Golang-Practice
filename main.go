package main

import "fmt"

func main() {
	var a [4]int
	fmt.Print(a)
	b := [4]int{1, 2, 3, 4}
	fmt.Println(b)
	c := [...]int{1, 2, 3, 4, 1, 2}
	fmt.Println(c)

	for index, value := range c {
		fmt.Println("index is ", index, "value is ", value)
	}

	data := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
	}
	fmt.Println(data)
	data1 := [...][]string{
		{"lion", "tiger", "there"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
	}
	fmt.Println(data1)
}
