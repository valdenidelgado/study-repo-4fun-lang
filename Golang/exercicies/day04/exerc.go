package main

import "fmt"

func main() {
	// //primeiro exercicio
	// arr := [5]int{1, 2, 3, 4, 5}

	// for _, v := range arr {
	// 	fmt.Println(v)
	// }
	// fmt.Printf("%T\n", arr)

	// //segundo exercicio
	// // slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// var slice []int

	// slice = append(slice, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	
	// for _, v := range slice {
	// 	fmt.Println(v)
	// }
		
	// fmt.Printf("%T\n", slice)

	// //terceiro exercicio
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	firstToThird := slice[:3]
	fmt.Println(firstToThird)

	fifthToLast := slice[4:]
	fmt.Println(fifthToLast)

	secondToSeventh := slice[1:7]
	fmt.Println(secondToSeventh)

	thirdToLastLessOne := slice[2:len(slice)-1]
	thirdToLastLessOne2 := slice[2:9]

	fmt.Println(thirdToLastLessOne)
	fmt.Println(thirdToLastLessOne2)
}