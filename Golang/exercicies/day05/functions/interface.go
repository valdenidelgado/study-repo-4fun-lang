package main

import "fmt"

type humano struct {
	nome string
	sobrenome string
}

type profissional struct {
	humano
	profissao string
}

type estudante struct {
	humano
	curso string
}

func (p profissional) oi() string {
	return "Olá, meu nome é " + p.nome + " e sou " + p.profissao
}

func (e estudante) oi() string {
	return "Olá, meu nome é " + e.nome + " e estou cursando " + e.curso
}

type ser interface {
	oi() string
}

// implementa a interface ser e referencia o tipo se é profissional ou estudante e retorna o metodo de cada um
func cumprimentar(s ser) {
	println(s.oi())
}

func main() {

	profissional1 := profissional{humano{"João", "Silva"}, "Programador"}
	estudante1 := estudante{humano{"Maria", "Silva"}, "Engenharia"}

	fmt.Println(profissional1.oi())
	fmt.Println(estudante1.oi())


	fmt.Println("====================================")

	cumprimentar(profissional1)
	cumprimentar(estudante1)
}