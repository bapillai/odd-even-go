package main

import "fmt"

var numList int

func main() {
	numList := [11]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range numList {
		if num%2 == 0 {
			fmt.Println(num, "is even")
		} else {
			fmt.Println(num, "is odd")
		}
	}
}
