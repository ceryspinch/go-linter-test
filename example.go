package main

import (
	"fmt"
)

const HELLO = "Hello!"

func hello_world() {
	fmt.Println("Hello World")
	fmt.Println("Hello World")
}

// This function prints the value of my_string which is "hey" on a new line
// ten times. It is called printTenTimes and this comment is very long
// blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah
// blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah
// blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah
// blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah
// blah blah blah blah blah blah blah blah blah blah blah blah blah blah blah
func printTenTimes() {
	my_string := "hey"

	fmt.Println(my_string)
	fmt.Println(my_string)
	fmt.Println(my_string)
	fmt.Println(my_string)
	fmt.Println(my_string)
	fmt.Println(my_string)
	fmt.Println(my_string)
	fmt.Println(my_string)
	fmt.Println(my_string)
	fmt.Println(my_string)

}

func checkX() bool {
	x := 0
	if x == 1 || x == 2 || x == 3 || x == 4 {
		return true
	}
	return false
}

func checkY() bool {
	x := 0
	if x%2 == 0 {
		if x == 2 || x == 4 {
			if x == 2 {
				return true
			}
			return true
		}
		return false
	}
	return false
}

func foo(a, b, c, d, e, f, g int) int {
	return a + b + c + d + e + f + g
}
