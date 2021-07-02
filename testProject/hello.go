package main

import "fmt"

var GSum int = 0

func SumTest(a int, b int) int {
	fmt.Printf("%d + %d = %d\n", a, b, a + b)
	GSum += a + b
	fmt.Printf("G_sum : %d\n", GSum)
	return a + b
}
