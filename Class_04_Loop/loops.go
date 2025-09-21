package main

import "fmt"

func main() {

	//  while loop
	// i := 0
	// for i<=3{
	// 	fmt.Println(" Iteration :",i)
	// 	i+=1
	// }

	// sum of n digits of number ;
	// var sum = 0
	// var n = 10
	// for 0 < n {
	// 	var temp = n % 10
	// 	sum += temp
	// 	n = n / 10
	// }
	// fmt.Println(sum)

	//  classic for loop
	//  for i:=0; i <=3;i++{
	// 	if(i ==2) {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	//  }

	//  range key word introduced in 1.2 version of go
	for i := range 3{
		fmt.Println(i)
	}


}