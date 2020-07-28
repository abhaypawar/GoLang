//If Else GoLang example and implementation with syntax
//Refer abhaypawar/GoLang repository for basic examples and projects build using GoLang

package main

import "fmt"

func main() {

	if 9%2 == 0 {
		fmt.Println("9 is even")
	} else {
		fmt.Println("9 is odd")
	}

	if 16%4 == 0 {
		fmt.Println("16 is divisible by 4")
	}

	if num := 3; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

/* The tentative output of the above program must be as below
9 is odd
16 is divisible by 4
3 has 1 digit
*/
