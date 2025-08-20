package main

import (
	"fmt"

	"github.com/barantoraman/DesignPatternsTraining/internal/creational/builder"
)

func main() {
	electricBuilder := builder.NewBuilder("electric")
	director := builder.NewDirector(electricBuilder)

	director.SetColor("black")
	director.SetEngine("honda vtec engine")
	director.SetWheel("default wheel")
	car := director.Build()

	fmt.Printf("car.Color: %v\n", car.Color)
	fmt.Printf("car.Engine: %v\n", car.Engine)
	fmt.Printf("car.Wheel: %v\n", car.Wheel)

}
