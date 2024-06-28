package main

import "fmt"

func genericSum[T float64 | int64](numbers ...T) T {
	var sum T = 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}
func main() {
	fmt.Println(genericSum(1.1, 1.1))
}
