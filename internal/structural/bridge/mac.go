package bridge

import "fmt"

type Mac struct {
	P Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.P.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.P = p
}
