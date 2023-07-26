package main

func main() {
	// example de variadic function

	// com isso a função rebece varios valores do tipo int sem precisar de varios parametros
	x := foo(1, 2, 3, 4, 5, 6, 7, 8, 9)
	println(x)

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// isso não funciona, pois a função espera varios valores do tipo int e não um slice
	// y := foo(slice)

	// para passar um slice para uma função variadic, é preciso desempacotar o slice
	// usando o operador ... (ellipsis) no final para desempacotar o slice
	y := foo(slice...)
	println(y)
}


func foo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}