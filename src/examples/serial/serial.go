package main

import (
	"machine"
	"time"
)

var (
	led = machine.LED
	d5  = machine.D5
	d6  = machine.D6
	d9  = machine.D9
	d10 = machine.D10
	d11 = machine.D11
	d12 = machine.D12
)

func main() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	d5.Configure(machine.PinConfig{Mode: machine.PinOutput})
	d6.Configure(machine.PinConfig{Mode: machine.PinOutput})
	d9.Configure(machine.PinConfig{Mode: machine.PinOutput})
	d10.Configure(machine.PinConfig{Mode: machine.PinOutput})
	d11.Configure(machine.PinConfig{Mode: machine.PinOutput})
	d12.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		led.Toggle()
		d5.High()
		println("hello world!hello world!hello world! hello world!")
		d5.Low()
		led.Low()
		time.Sleep(100 * time.Millisecond)
	}
}
