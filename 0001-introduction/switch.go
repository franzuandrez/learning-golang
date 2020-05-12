package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Println("Wirte ", i, " as")

	switch i {

	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It´s the weekend")
	default:
		fmt.Println("It's a Weekday")

	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Its before noon")
	default:
		fmt.Println("Its after noon")
	}

	WhatAmI := func(i interface{}) {

		switch t := i.(type) {

		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("Im an INT")
		default:
			fmt.Println("IDK", t)
		}
	}

	WhatAmI(true)
	WhatAmI(1)
	WhatAmI("Franzua")
}
