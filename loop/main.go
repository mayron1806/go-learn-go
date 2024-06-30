package main

import "fmt"

func forDefault() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
func interateArray() {
	names := []string{"a", "b", "c"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}
func interateArray2() {
	names := []string{"a", "b", "c"}
	for index, name := range names {
		fmt.Println(index, name)
	}
}
func while() {
	names := []string{"a", "b", "c"}
	var i int
	for i < len(names) {
		fmt.Println(names[i])
		i++
	}
}
func main() {
	// forDefault()
	// interateArray()
	// interateArray2()
	while()
}
