// package main

// import (
// 	"fmt"
// 	"time"
// )

// func showNumbers(done chan<- bool) {
// 	for n := 1; n < 24; n++ {
// 		fmt.Printf("%d", n)
// 		time.Sleep(time.Millisecond * 100)
// 	}
// 	done <- true
// }
// func showLetters(done chan<- bool) {
// 	for l := 'a'; l <= 'z'; l++ {
// 		fmt.Printf("%c", l)
// 		time.Sleep(time.Millisecond * 100)
// 	}
// 	done <- true
// }

// func main() {
// 	cNumbers := make(chan bool) // cria channel
// 	cLetters := make(chan bool) // cria channel
// 	go showLetters(cLetters)
// 	go showNumbers(cNumbers)
// 	<-cLetters
// 	<-cNumbers
// }

// channel with buffer
package main

import (
	"fmt"
	"time"
)

func showNumbers(cNumber chan<- int) {
	for n := 1; n <= 10; n++ {
		cNumber <- n
		fmt.Printf("writed: %d\n", n)
		time.Sleep(time.Millisecond * 100)
	}
	close(cNumber)
}

func main() {
	cNumbers := make(chan int, 3) // cria channel

	go showNumbers(cNumbers)

	for v := range cNumbers {
		fmt.Printf("value: %d\n", v)
		time.Sleep(time.Millisecond * 1000)
	}

}
