package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	person := make(map[string]string)

	var name, address string

	fmt.Printf("Please Enter Your name: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		name = scanner.Text()
	}

	fmt.Printf("Please Enter Your Address: ")
	if scanner.Scan() {
		address = scanner.Text()
	}

	person["name"] = name
	person["address"] = address

	obj, err := json.Marshal(person)

	if err != nil {
		fmt.Println("error")
	}

	fmt.Printf("%s\n", obj)
}


