package main

import "fmt"

type veiculo struct {
	portas int
	cor string
}

type caminhonete struct {
	veiculo
	tracaoNasQuatro bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {

	caminhonete1 := caminhonete{veiculo{4, "preta"}, true}
	sedan1 := sedan{veiculo{4, "branca"}, false}

	fmt.Println(caminhonete1)
	fmt.Println(sedan1)

	fmt.Println(caminhonete1.portas)
	fmt.Println(sedan1.cor)
}