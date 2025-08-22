package main

import (
	"github.com/barantoraman/DesignPatterns/internal/behavioral/cor"
)

func main() {
	cashier := &cor.Cashier{}

	medical := &cor.Medical{}
	medical.SetNext(cashier)

	doc := &cor.Doctor{}
	doc.SetNext(medical)

	reception := &cor.Reception{}
	reception.SetNext(doc)

	patient := &cor.Patient{Name: "abc"}
	reception.Execute(patient)
}
