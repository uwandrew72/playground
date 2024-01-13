package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println(a)
	fmt.Printf("type of a = %T\n", a)
	var b int = 100
	fmt.Println(b)
	var c = 100
	fmt.Println(c)
	fmt.Printf("type of c = %T\n", c)

	var s1 string = "abcd"
	fmt.Printf("type of s1 = %T\n", s1)

	e := 100
	fmt.Printf("type of e = %T\n", e)
}
