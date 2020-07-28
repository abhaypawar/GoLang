//for loop GoLang example and implementation with syntax
//Refer abhaypawar/GoLang repository for basic examples and projects build using GoLang

package main

import "fmt"

func main() {

	// ForType1 (most basic type), with single condition//
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// ForType2 (Initial/Condition/After) for loop//
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// ForType3 (Type without a condition) will loop repeatedly unless break condition is applied//
	for {
		fmt.Println("loop")
		break
	}

	// ForType4 (continue keyword to the next iteration of the loop//
	for n := 0; n <= 50; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

//the tentatite output should be :
/*1
2
3
7
8
9
loop
1
3
5
7
9
11
13
15
17
19
21
23
25
27
29
31
33
35
37
39
41
43
45
47
49
*/
