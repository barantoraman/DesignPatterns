package main

import (
	"fmt"

	simplefactory "github.com/barantoraman/DesignPatternsTraining/internal/creational/simple_factory"
)

func main() {
	ak47 := simplefactory.GetGun("ak47")
	ak47.SetName("ak47")
	ak47.SetPower(7)

	fmt.Printf("ak47.GetPower(): %v\n", ak47.GetPower())
	fmt.Printf("ak47.GetName(): %v\n", ak47.GetName())
}
