package main

import (
	"fmt"
	"reflect"
)


	/*
	   Data types specify the type of data that a valid Go variable can hold. In Go language, the type is divided into four categories which are as follows:

	   Basic type: Numbers, strings, and booleans come under this category.

	   Aggregate type: Array and structs come under this category.

	   Reference type: Pointers, slices, maps, functions, and channels come under this category.

	   Interface type: Interfaces in Go define a set of method signatures that a type must implement. They allow you to write flexible and reusable code by focusing on behavior rather than specific implementation. */
func main() {
	var a = 882
	fmt.Println("String :", reflect.TypeOf("Hello world"))
	fmt.Println("Integer :", reflect.ValueOf(a).Kind())
	fmt.Println("Boolean :", reflect.TypeOf(true))

}
