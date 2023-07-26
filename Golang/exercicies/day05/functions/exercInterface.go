package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	lado float64
}

type circulo struct {
	raio float64
}

func (q quadrado) area() float64 {
	return q.lado * q.lado
}

func (c circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}

type forma interface {
	area() float64
}

func info(f forma) float64 {
	return f.area()
}

func main() {

	q := quadrado{15}
	c := circulo{5.25}

	fmt.Println(info(q))
	fmt.Println(info(c))
}