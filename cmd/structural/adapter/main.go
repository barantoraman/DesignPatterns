package main

import "github.com/barantoraman/DesignPatterns/internal/structural/adapter"

func main() {
	mac := &adapter.Mac{}
	windows := &adapter.Windows{}
	winAdapter := &adapter.WindowsAdapter{
		WindowsMachine: windows,
	}

	mac.InsertIntoLightningPort()
	windows.InsertIntoUSBPort()
	winAdapter.InsertIntoLightningPort()
}
