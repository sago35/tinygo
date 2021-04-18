package main

import (
	"device/sam"
	"fmt"
	"machine"
	"time"
)

var led = machine.LED

func main() {
	stb := machine.CAN1_STANDBY
	stb.Configure(machine.PinConfig{Mode: machine.PinOutput})
	stb.Low()

	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	can1 := machine.CAN{Bus: sam.CAN1}
	can1.Configure(machine.CANConfig{
		Frequency: 500000,
		Rx:        machine.CAN1_RX,
		Tx:        machine.CAN1_TX,
	})

	can0 := machine.CAN{Bus: sam.CAN0}
	can0.Configure(machine.CANConfig{
		Frequency: 500000,
		Rx:        machine.CAN0_RX,
		Tx:        machine.CAN0_TX,
	})

	for {
		led.Toggle()

		can1.Tx(0x123, 8, []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF})
		time.Sleep(time.Microsecond * 300)
		can1.TxFD(0x789, 16, []byte{0x02, 0x24, 0x46, 0x67, 0x89, 0xAB, 0xCD, 0xEF, 0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF})
		time.Sleep(time.Millisecond * 1000)

		sz0 := can0.Size()
		if sz0 > 0 {
			fmt.Printf("CAN0 %d\r\n", sz0)
			for i := 0; i < sz0; i++ {
				id, dlc, data := can0.Rx()
				fmt.Printf("-> %08X %X %#v\r\n", id, dlc, data)
			}
		}

		sz1 := can1.Size()
		if sz1 > 0 {
			fmt.Printf("CAN1 %d\r\n", sz1)
			for i := 0; i < sz1; i++ {
				id, dlc, data := can1.Rx()
				fmt.Printf("-> %08X %X %#v\r\n", id, dlc, data)
			}
		}
	}
}
