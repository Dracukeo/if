package main

import "fmt"

func main() {
	x := 1

	if x > 6 {
		fmt.Print("More than 6")
	} else if x < 2 {
		fmt.Print("Less than 2")
	} else {
		fmt.Print("Whatever.. Something Else.")
	}
}
