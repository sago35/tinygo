package main

import (
	"device/sam"
	"fmt"
	"machine"
	"time"
)

var led = machine.LED

func main() {
	// CAN1_STB : PC13
	// CAN1_RX  : PB13
	// CAN1_TX  : PB12

	vcom := machine.UART4
	vcom.Configure(machine.UARTConfig{TX: machine.UART4_TX_PIN, RX: machine.UART4_RX_PIN})
	//fmt.Fprintf(vcom, "115200 bps で接続すれば Virtual COM Port も使えます\r\n")

	// Activate CANtr
	stb := machine.PC13
	stb.Configure(machine.PinConfig{Mode: machine.PinOutput})

	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	can1 := machine.CAN{Bus: sam.CAN1}
	can1.Configure(machine.CANConfig{
		Frequency: 500000,
		Rx:        machine.PB13,
		Tx:        machine.PB12,
	})

	can0 := machine.CAN{Bus: sam.CAN0}
	can0.Configure(machine.CANConfig{
		Frequency: 500000,
		Rx:        machine.PA23,
		Tx:        machine.PA22,
	})

	for {
		led.Toggle()
		can1.Tx(0x123, 8, []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF})
		//can1.TxFD(0x789, 16, []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF, 0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF})
		//can0.Tx(0x456, 8, []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF})
		time.Sleep(time.Microsecond * 300)
		can1.TxFD(0x789, 16, []byte{0x02, 0x24, 0x46, 0x67, 0x89, 0xAB, 0xCD, 0xEF, 0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF})
		time.Sleep(time.Millisecond * 1000)

		sz := can0.Size()
		fmt.Fprintf(vcom, "%#d\r\n", sz)
		for i := 0; i < sz; i++ {
			id, dlc, data := can0.Rx()
			fmt.Fprintf(vcom, "-> %08X %X %#v\r\n", id, dlc, data)
		}
	}
}
