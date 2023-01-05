package main

import (
	"fmt"
	"learning/interface/shape"
)

func main() {
	circle := shape.NewCircle(2)
	fmt.Println(getSquareArea(circle))
	// fmt.Println(circle.Area())

	square := shape.NewSquare(4)
	fmt.Println(getSquareArea(square))
	// fmt.Println(square.Area())

	// shapes := []shape.Figure{
	// 	circle,
	// 	square,
	// }

	// for _, s := range shapes{
	// 	fmt.Println(s.Area())
	// }

}

func getSquareArea(s shape.Figure) float64 {
	area := s.Area()

	return area
}