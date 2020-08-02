package main

import "fmt"

func dmain() {
	fmt.Println("Enter the floating point number to be truncated to int")
	var f float64
	fmt.Scan(&f)
	var i int = int(f)
	fmt.Println("the floating point number that is  truncated to int is ", i)

}
