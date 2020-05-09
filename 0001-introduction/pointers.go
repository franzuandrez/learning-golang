package main

import "fmt"

func main() {
	x := 15
	a := &x //Memory access

	fmt.Println(a)
	fmt.Println(*a)
	*a = 5
	fmt.Println(x)

}
