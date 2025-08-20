package bridge

import "fmt"

type Computer interface {
	Print()
	SetPrinter(Printer)
}

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

type Printer interface {
	PrintFile()
}

type Epson struct{}

func (p *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}

type Hp struct {
}

func (p *Hp) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}
