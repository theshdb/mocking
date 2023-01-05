package shape

//go:generate mockgen -destination=../mocks/mock_figure.go -package=mocks . Figure
import (
	"fmt"
	"math"
)

type Figure interface{
	Area() float64
}

type Circle struct{
	Radius float64
}

type Square struct{
	Length float64
}

func (s Square) Area() float64 {
    return s.Length * s.Length
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func NewSquare(s float64) Figure{
	return Square{Length: s}
}

func NewCircle(c float64) Figure{
	return Circle{Radius: c}
}

func getArea(shape Figure) {

    fmt.Println(shape.Area())
}

