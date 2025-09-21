package main

import "fmt"

func main() {

	// Numbers
	// var a int = 64
	// fmt.Println("SIgned INterger :", a)
	// var af float32 = 64.5
	// fmt.Println("SIgned Float :", af)
	// //  Booleans
	// var b bool = false
	// fmt.Println("Boolean value :", b)
	// //  Strings
	// var str string = "Hello World"
	// fmt.Println("Boolean value :", str)

	//  CONSTANTS -> once declared it never changed

	const name = "go lang " // it never changed again

	const (
		port = 8080
		host = "localHost"
	)
	fmt.Println("Constant Name :", name)
	fmt.Printf("Constant port %d and host is %s :", port, host)

}
