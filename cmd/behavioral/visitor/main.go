package main

import (
	"fmt"

	"github.com/barantoraman/DesignPatternsTraining/internal/behavioral/visitor"
)

func main() {
	square := &visitor.Square{Side: 2}
	circle := &visitor.Circle{Radius: 3}
	rectangle := &visitor.Rectangle{L: 2, B: 3}

	areaCalculator := &visitor.AreaCalculator{}

	square.Accept(areaCalculator)
	circle.Accept(areaCalculator)
	rectangle.Accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &visitor.MiddleCoordinates{}
	square.Accept(middleCoordinates)
	circle.Accept(middleCoordinates)
	rectangle.Accept(middleCoordinates)
}
