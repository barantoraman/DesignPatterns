package adapter

import "fmt"

type USBPort interface {
	InsertIntoUSBPort()
}

type Windows struct{}

type WindowsAdapter struct {
	WindowsMachine USBPort
}

func (w *Windows) InsertIntoUSBPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

func (a *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	a.WindowsMachine.InsertIntoUSBPort()
}
