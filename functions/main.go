package main

import "fmt"

func main() {
	result := sum(1, 2)
	fmt.Println(result)

	fmt.Println(multiReturnValues(1, 5))

	fmt.Println(multiReturnValues2(199, 213))
	v1, v2, res := multiReturnValues(22, 2)
	fmt.Println(v1, v2, res)
}

func sum(a, b int) int {
	result := a + b
	return result
}
func multiReturnValues(a, b int) (num1, num2, result int) {
	num1 = a
	num2 = b
	result = a + b
	return
}
func multiReturnValues2(a int, b int) (string, int) {
	str := fmt.Sprintf("A soma de %v e %v é: %v", a, b, a+b)
	return str, a + b
}
