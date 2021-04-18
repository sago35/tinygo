// +build feather_m4_can

package main

import (
	"machine"
)

func init() {
	// power on the CAN Transceiver
	boost_en := machine.BOOST_EN
	boost_en.Configure(machine.PinConfig{Mode: machine.PinOutput})
	boost_en.High()
}
