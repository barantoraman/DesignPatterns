package main

import (
	"fmt"
	"log"

	"github.com/barantoraman/DesignPatterns/internal/behavioral/state"
)

func main() {
	vendingMachine := state.NewVendingMachine(1, 10)

	err := vendingMachine.RequestItem()
	if err != nil {
		log.Fatalf("err: %s", err.Error())
	}

	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalf("err: %s", err.Error())
	}

	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatalf("err: %s", err.Error())
	}

	fmt.Println()

	err = vendingMachine.AddItem(2)
	if err != nil {
		log.Fatalf("err: %s", err.Error())
	}

	fmt.Println()

	err = vendingMachine.RequestItem()
	if err != nil {
		log.Fatalf("err: %s", err.Error())
	}

	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalf("err: %s", err.Error())
	}

	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatalf("err: %s", err.Error())
	}
}
