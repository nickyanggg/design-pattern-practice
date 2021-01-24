package main

func main() {
	client := &client{}
	mac := &mac{}
	client.insertSquareUsbInComputer(mac)
	windowsAdapter := &windowsAdapter{
		windowMachine: &windows{},
	}
	client.insertSquareUsbInComputer(windowsAdapter)
}
