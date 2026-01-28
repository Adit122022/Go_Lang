//  use SWITCH CASE if more  conditions u need to check

package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Print("Enter the day number :- ")
	fmt.Scanf("%v", &i)

	//  first way to use switch
	// switch i {
	// case 0:
	// 	fmt.Println("Sunday")
	// case 1:
	// 	fmt.Println("Monday")
	// case 2:
	// 	fmt.Println("Tuesday")
	// case 3:
	// 	fmt.Println("Wednesday")
	// case 4:
	// 	fmt.Println("Thrusday")
	// case 5:
	// 	fmt.Println("Friday")
	// case 6:
	// 	fmt.Println("Saturday")
	// default: // defualt is not neccesary in go
	// 	fmt.Println(" DEFAULT")
	// }

	//  second way to use switch
	// switch time.Now().Weekday() {
	// case time.Sunday, time.Saturday:
	// 	fmt.Println("It Is A WEEkEND");
	// default:
	// 	fmt.Println(("It is not A weekday"))
	// }

	//  type switch  (function )
	whoAmi := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Println("I AM INTEGER")
		case string:
			fmt.Println("I AM STRING")
		case bool:
			fmt.Println("I AM BOOLEAN")
		case float64:
			fmt.Println("I AM FLOAT")
		case nil:
			fmt.Println("I AM NIL")
		default:
			fmt.Println("I AM SOMETHING ELSE", v)
		}
	}
	// Example calls to whoAmi to use the function and avoid "i declared and not used" error
	whoAmi(i)
}
