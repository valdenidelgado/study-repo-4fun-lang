package main

import "fmt"

func main() {
    fmt.Println(Greetings("World"))
}

func Greetings(name string) string {
	return "Hello, " + name + "!"
}