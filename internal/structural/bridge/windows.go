package bridge

import "fmt"

type Windows struct {
	P Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.P.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.P = p
}
