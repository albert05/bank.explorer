package main

import (
	"fmt"
)

func main() {

}


func goTest() {
	l := []int{1,2,3,4,5}
	ch := make(chan bool)

	for _, i := range l {
		go test(i, ch)
	}

	for _, _ = range l {
		<-ch
	}

	close(ch)
}

func test(i int, ch chan<- bool) {
	fmt.Println(i)
	ch <- true
}