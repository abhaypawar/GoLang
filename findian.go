package main

import (
	"fmt"
	"strings"
)

func main() {

	str := ""
	fmt.Printf("Enter String: ")
	fmt.Scan(&str)
	var x int = len(str)
	if str[0]=='i' && str[x-1]=='n' && strings.Contains(str,"a") { 
            		fmt.Printf("Found") 
	}    else 	{
		fmt.Printf("Not Found") 	
		}
}