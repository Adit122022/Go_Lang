//  use SWITCH CASE if more  conditions u need to check

package main

import "fmt"

func main() {
	var i int
	fmt.Print("Enter the day number :- ")
	fmt.Scanf("%v", &i)
	switch i {
	case 0:
		fmt.Println("Sunday")
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thrusday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	default:
		fmt.Println(" DEFAULT")
	}
}
