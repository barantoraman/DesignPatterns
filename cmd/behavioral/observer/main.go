package main

import "github.com/barantoraman/DesignPatterns/internal/behavioral/observer"

func main() {
	shirtItem := observer.NewItem("nike shirt")
	observerFirst := &observer.Customer{ID: "test123@mail.com"}
	observerSecond := &observer.Customer{ID: "test456@mail.com"}

	shirtItem.Register(observerSecond)
	shirtItem.Register(observerFirst)
	shirtItem.UpdateAvailability()
}
