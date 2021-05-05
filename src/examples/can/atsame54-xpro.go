// +build atsame54_xpro

package main

import (
	"machine"
)

func init() {
	// Activate CAN Transceiver
	stb := machine.CAN1_STANDBY
	stb.Configure(machine.PinConfig{Mode: machine.PinOutput})
	stb.Low()
}
