package main

import (
	"device/sam"
	"fmt"
	"machine"
	"runtime/interrupt"
	"time"
)

var led = machine.LED

func main() {
	// CAN1_STB : PC13
	// CAN1_RX  : PB13
	// CAN1_TX  : PB12

	pb04 := machine.PB04
	pb04.Configure(machine.PinConfig{Mode: machine.PinOutput})

	vcom := machine.UART4
	vcom.Configure(machine.UARTConfig{TX: machine.UART4_TX_PIN, RX: machine.UART4_RX_PIN})
	//fmt.Fprintf(vcom, "115200 bps で接続すれば Virtual COM Port も使えます\r\n")

	fmt.Fprintf(vcom, "..\r\n")
	//fmt.Fprintf(vcom, "%p\r\n", machine.HandleCANIntr)

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

	//machine.DumpCanCallback()
	//fmt.Fprintf(vcom, "%#v\r\n", machine.CanCallbacks)
	//machine.SetCANInterrupt()
	//fmt.Fprintln(vcom, "rx\r\n")
	//time.Sleep(1 * time.Second)

	sam.CAN0.ILE.SetBits(sam.CAN_ILE_EINT0)
	sam.CAN0.IE.SetBits(sam.CAN_IE_RF0NE)
	interrupt.New(sam.IRQ_CAN0, func(interrupt.Interrupt) {
		sam.CAN0.IR.SetBits(sam.CAN_IR_RF0N)
		pb04.Toggle()
	}).Enable()

	//interrupt.New(sam.IRQ_CAN0, func(interrupt.Interrupt) {}).Enable()
	//fmt.Fprintf(vcom, "%#v\r\n", machine.CanCallbacks)

	for {
		if true {
			led.Toggle()
		}
		can1.Tx(0x123, 8, []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF})
		//can1.TxFD(0x789, 16, []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF, 0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF})
		//can0.Tx(0x456, 8, []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF})
		//	time.Sleep(time.Microsecond * 300)
		//	can1.TxFD(0x789, 16, []byte{0x02, 0x24, 0x46, 0x67, 0x89, 0xAB, 0xCD, 0xEF, 0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF})

		fmt.Fprintf(vcom, "sam.CAN_IR : %08X\r\n", can0.Bus.IR.Get())
		time.Sleep(time.Millisecond * 1000)
		fmt.Fprintf(vcom, "sam.CAN_IR : %08X\r\n", can0.Bus.IR.Get())

		sz := can0.Size()
		fmt.Fprintf(vcom, "%#d\r\n", sz)
		for i := 0; i < sz; i++ {
			id, dlc, data := can0.Rx()
			if false {
				fmt.Fprintf(vcom, "-> %08X %X %#v\r\n", id, dlc, data)
			}
		}
	}
}
