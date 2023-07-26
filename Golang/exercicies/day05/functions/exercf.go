package main

import "fmt"

func main() {
	func (x int) {
		fmt.Println(x)
	}(10)

	x := func (x int) int {
		return x * 2
	}

	fmt.Println(x(10))

	x = returnFunc()

	fmt.Println(x(10))
	foo(printHello)
	closure("João")()
}

func returnFunc() func(int) int {
	return func (x int) int {
		return x * 2
	}
}

func printHello() {
	fmt.Println("PRESTOU TENÇÃO EM")
}


func foo(f func()) {
	fmt.Println("prestatenção")
	f()
}

func closure(name string) func() {
	sobrenome := "FALTOU SOBRENOME"
	return func() {
		fmt.Println("Hello", name, sobrenome)
	}
}