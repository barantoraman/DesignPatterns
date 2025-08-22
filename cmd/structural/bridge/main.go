package main

import "github.com/barantoraman/DesignPatterns/internal/structural/bridge"

func main() {
	hpPrinter := &bridge.Hp{}
	epsonPrinter := &bridge.Epson{}

	mac := &bridge.Mac{}
	mac.SetPrinter(epsonPrinter)
	mac.Print()

	mac.SetPrinter(hpPrinter)
	mac.Print()
}
