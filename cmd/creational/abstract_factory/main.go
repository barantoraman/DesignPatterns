package main

import (
	"fmt"
	"log"

	abstractfactory "github.com/barantoraman/DesignPatterns/internal/creational/abstract_factory"
)

func main() {
	factory, err := abstractfactory.GetFurnitureFactory("victorian")
	if err != nil {
		log.Fatal("failed to get furniture factory: %w", err)
	}

	chair := factory.MakeChair()
	sofa := factory.MakeSofa()

	chair.SetColor("red")
	sofa.SetColor("redd")

	fmt.Printf("chair.GetColor(): %v\n", chair.GetColor())
	fmt.Printf("sofa.GetColor(): %v\n", sofa.GetColor())
}
