package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type person struct {
	fname string
	lname string
}

func main() {
	fmt.Printf("Enter the name of the File: ")
	console := bufio.NewReader(os.Stdin)
	filename, _ := console.ReadString('\n')

	file, err := os.Open(strings.TrimSpace(filename))

	sli := make([]person, 0, 5)

	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			name := strings.Fields(strings.TrimSpace(scanner.Text()))
			var p person
			p.fname = name[0]
			p.lname = name[1]
			sli = append(sli, p)
		}
                for idx,dat := range sli{
			fmt.Print(idx,"First Name: ",dat.fname," Last Name: ",dat.lname)
                        fmt.Printf("\n")
                }
		fmt.Print(sli)
		fmt.Printf("\n")
	}
}


