package main

import (
	"fmt"
)

var n int
var divisible int

func main() {

	for {

		divisible = 0
		fmt.Scanln(&n)

		for i := 2; i <= n/2; i++ {

			if n%i == 0 {

				divisible++

			}
		}

		if divisible == 0 {

			fmt.Printf("o %d primo", n)

		}

	}

}
