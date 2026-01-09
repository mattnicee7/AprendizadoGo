package formas

import (
	"fmt"
	"math"
)

type Forma interface {
	Area() float64
}

func EscreverArea(f Forma) {
	fmt.Printf("area: %.02f\n", f.Area())
}

type Retangulo struct {
	Altura  float64
	Largura float64
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}

func main() {
	r := Retangulo{10, 15}
	EscreverArea(r)

	c := Circulo{10}
	EscreverArea(c)
}
