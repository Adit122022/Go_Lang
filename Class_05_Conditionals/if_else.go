package main

import "fmt"

func main() {
	age := 15
	if age >= 18 {
		fmt.Println("Person is an Adult  , AGE :- ", age)
	} else if age >= 12 {
		fmt.Println("Person is teenager , AGE :-", age)
	} else {
		fmt.Println("Person is under age  !!  AGE :- ", age)
	}

	// go does not have ternary operator , you will have to use if else
}
