package main

import (
	"fmt"
	"time"
)

func main() {
	go foo()
	go foo()
	foo()

}

func foo() {
	for i := 0; i < 20; i++ {
		fmt.Println("foo:", i)
		time.Sleep(time.Second)
	}
}
