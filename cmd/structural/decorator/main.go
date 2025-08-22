package main

import (
	"fmt"

	"github.com/barantoraman/DesignPatterns/internal/structural/decorator"
)

func main() {
	pizza := &decorator.VeggieMania{}
	pizzaWithCheese := &decorator.CheeseTopping{P: pizza}
	pizzaWithCheeseTomato := &decorator.TomatoTopping{P: pizzaWithCheese}
	fmt.Printf("pizzaWithCheeseTomato.GetPrice(): %v\n", pizzaWithCheeseTomato.GetPrice())

}
