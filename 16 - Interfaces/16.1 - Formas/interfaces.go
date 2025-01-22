package main

import (
	"fmt"
	"math"
)

type forma interface{
	area () float64
}

type retangulo struct{
	altura float64
	largura float64
}

type circulo struct {
	raio float64
}

func escreverArea(f forma) {
	fmt.Printf("A area da forma é de %0.2f\n", f.area())
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {
	fmt.Println("interfaces")
	r := retangulo{20, 40}
	c := circulo{20}
	escreverArea(r)
	escreverArea(c)
}