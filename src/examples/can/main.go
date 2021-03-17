package main

import (
	"machine"
	"time"
)

var led = machine.LED

func main() {
	// CAN1_STB : PC13
	// CAN1_RX  : PB13
	// CAN1_TX  : PB12
	//time.Sleep(3 * time.Second)
	//fmt.Printf("\r\n")

	// Activate CANtr
	stb := machine.PC13
	stb.Configure(machine.PinConfig{Mode: machine.PinOutput})

	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	machine.InitCAN()
	for {
		led.Toggle()
		machine.Tx(0x123, 8, []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF})
		//TxFD(0x123, 12, []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF, 0xA1, 0xA2, 0xA3, 0xA4})
		time.Sleep(time.Millisecond * 1000)
	}
}
