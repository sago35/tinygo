// +build sam,atsamd51,atsamd51g19

// Peripheral abstraction layer for the atsamd51.
//
// Datasheet:
// http://ww1.microchip.com/downloads/en/DeviceDoc/60001507C.pdf
//
package machine

import "device/sam"

const HSRAM_SIZE = 0x00030000

// InitPWM initializes the PWM interface.
func InitPWM() {
	// turn on timer clocks used for PWM
	sam.MCLK.APBBMASK.SetBits(sam.MCLK_APBBMASK_TCC0_ | sam.MCLK_APBBMASK_TCC1_)
	sam.MCLK.APBCMASK.SetBits(sam.MCLK_APBCMASK_TCC2_)

	//use clock generator 0
	sam.GCLK.PCHCTRL[sam.PCHCTRL_GCLK_TCC0].Set((sam.GCLK_PCHCTRL_GEN_GCLK0 << sam.GCLK_PCHCTRL_GEN_Pos) |
		sam.GCLK_PCHCTRL_CHEN)
	sam.GCLK.PCHCTRL[sam.PCHCTRL_GCLK_TCC2].Set((sam.GCLK_PCHCTRL_GEN_GCLK0 << sam.GCLK_PCHCTRL_GEN_Pos) |
		sam.GCLK_PCHCTRL_CHEN)
}

// getTimer returns the timer to be used for PWM on this pin
func (pwm PWM) getTimer() *sam.TCC_Type {
	switch pwm.Pin {
	case PA08:
		return sam.TCC0
	case PA09:
		return sam.TCC0
	case PA10:
		return sam.TCC0
	case PA11:
		return sam.TCC0
	case PA12:
		return sam.TCC0
	case PA13:
		return sam.TCC0
	case PA14:
		return sam.TCC2
	case PA15:
		return sam.TCC2
	case PA16:
		return sam.TCC1
	case PA17:
		return sam.TCC1
	case PA18:
		return sam.TCC1
	case PA19:
		return sam.TCC1
	case PA20:
		return sam.TCC1
	case PA21:
		return sam.TCC1
	case PA22:
		return sam.TCC1
	case PA23:
		return sam.TCC1
	case PA24:
		return sam.TCC2
	case PA30:
		return sam.TCC2
	case PA31:
		return sam.TCC2

	case PB02:
		return sam.TCC2
	case PB10:
		return sam.TCC0
	case PB11:
		return sam.TCC0

	default:
		return nil // not supported on this pin
	}
}
