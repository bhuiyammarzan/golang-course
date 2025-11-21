package main

import (
	"fmt"
	"time"
)

func main () {
	// simple switch
	// i := 3

	// switch i {
	// case 1:
	// 	fmt.Println("one")
	
	// case 2: 
	// 	fmt.Println("two")

	// case 3: 
	// 	fmt.Println("three")
	
	// default:
	// 	fmt.Println("others")
	// }


	// multiple condition switch
	switch time.Now().Weekday(){
	case time.Friday, time.Saturday:
		fmt.Println("it's weekend")

	default:
		fmt.Println("it's workday")
	}

	whoAmI := func (i interface{})  {
		
		switch t := i.(type) {
		case int:
			fmt.Println("it's an inter")

		case string: 
			fmt.Println("it's an string")

		case bool: 
			fmt.Println("it's an boolean")

		default:
			fmt.Println("other", t)
		}

	}

	whoAmI(40.5)
}