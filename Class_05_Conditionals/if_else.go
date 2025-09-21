package main

import "fmt"

func main() {
	// age := 12
	// if age >= 18 {
	// 	fmt.Println("Person is adult")
	// } else if age >= 12 && age < 18 {
	// 	fmt.Println("Person is  Teenager")
	// } else {
	// 	fmt.Println("Person is Under 18")
	// }


	//  another example 
	//  role := "admin"
	//  hasPermisson :=true
	//   if(role == "admin" || hasPermisson){
	// 	fmt.Println("Show Admin Dashboard")
	//   }else {
	// 	fmt.Println("Show User Dashboard")
	//   }

	//  directly declare  varibale in if construct

	if age := 15;age >= 18 {
		fmt.Println("Person is an Adult  , AGE :- ", age)
	}else if age >=12{
		fmt.Println("Person is teenager , AGE :-", age)
	} else {
		fmt.Println("Person is under age  !!  AGE :- " , age)
	 }

	//   go does not have ternary operator , you will have to use if else
	 }
