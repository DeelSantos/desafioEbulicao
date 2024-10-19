package main

import "fmt"

const ebulicaoF = 273
const ebulicaoK = 373

var ebulicaoC int

func main() {

	ebulicaoC = ebulicaoK - ebulicaoF

	fmt.Printf("O ponto de ebulição de Kelvin para Celsius é: %d ºC", ebulicaoC)
}
