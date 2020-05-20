// +build sam,atsamd51,wioterminal

package machine

import "device/sam"

// used to reset into bootloader
const RESET_MAGIC_VALUE = 0xf01669ef

// GPIO Pins
const (
	D0  = PB17 // UART0 RX/PWM available
	D1  = PB16 // UART0 TX/PWM available
	D4  = PA14 // PWM available
	D5  = PA16 // PWM available
	D6  = PA18 // PWM available
	D8  = PB03 // built-in neopixel
	D9  = PA19 // PWM available
	D10 = PA20 // can be used for PWM or UART1 TX
	D11 = PA21 // can be used for PWM or UART1 RX
	D12 = PA22 // PWM available
	D13 = PA23 // PWM available
	D21 = PA13 // PWM available
	D22 = PA12 // PWM available
	D23 = PB22 // PWM available
	D24 = PB23 // PWM available
	D25 = PA17 // PWM available
)

// Analog pins
const (
	A0 = PB08 // ADC/AIN[0]
	A1 = PB09 // ADC/AIN[2]
	A2 = PA07 // ADC/AIN[3]
	A3 = PB04 // ADC/AIN[4]
	A4 = PB05 // ADC/AIN[5]
	A5 = PB06 // ADC/AIN[10]
	A6 = PB04 // ADC/AIN[10]
	A7 = PB07 // ADC/AIN[10]
	A8 = PA06 // ADC/AIN[10]

	ADC0 = A0
	ADC1 = A1
	ADC2 = A2
	ADC3 = A3
	ADC4 = A4
	ADC5 = A5
	ADC6 = A6
	ADC7 = A7
	ADC8 = A8
)

const (
	LED = PA15

	BUTTON_1  = PC26
	BUTTON_2  = PC27
	BUTTON_3  = PC28
	WIO_KEY_A = BUTTON_1
	WIO_KEY_B = BUTTON_2
	WIO_KEY_C = BUTTON_3

	BUTTON = BUTTON_1

	SWITCH_X = PD20
	SWITCH_Y = PD12
	SWITCH_Z = PD09
	SWITCH_B = PD08
	SWITCH_U = PD10

	WIO_5S_UP    = SWITCH_X
	WIO_5S_LEFT  = SWITCH_Y
	WIO_5S_RIGHT = SWITCH_Z
	WIO_5S_DOWN  = SWITCH_B
	WIO_5S_PRESS = SWITCH_U

	OUTPUT_CTR_5V  = PC14
	OUTPUT_CTR_3V3 = PC15

	LCD_MISO = PB18
	LCD_MOSI = PB19
	LCD_SCK  = PB20
	LCD_CS   = PB21

	LCD_BACKLIGHT_CTR = PC05
	LCD_DC            = PC06
	LCD_RESET         = PC07

	LCD_XL = PC10
	LCD_YU = PC11
	LCD_XR = PC12
	LCD_YD = PC13
)

// UART0 aka USBCDC pins
const (
	USBCDC_DM_PIN = PA24
	USBCDC_DP_PIN = PA25
)

// UART1 pins
const (
	UART_TX_PIN = D1
	UART_RX_PIN = D0
)

// UART2 pins
const (
	UART2_TX_PIN = A4
	UART2_RX_PIN = A5
)

// I2C pins
const (
	SDA_PIN = D21 // SDA: SERCOM4/PAD[0]
	SCL_PIN = D22 // SCL: SERCOM4/PAD[1]
)

// I2C on the Feather M4.
var (
	I2C0 = I2C{
		Bus:    sam.SERCOM4_I2CM,
		SERCOM: 4,
	}
)

// SPI pins
const (
	SPI0_SCK_PIN  = PB03 // SCK: SERCOM5/PAD[1]
	SPI0_MOSI_PIN = PB02 // MOSI: SERCOM5/PAD[0]
	SPI0_MISO_PIN = PB00 // MISO: SERCOM5/PAD[2]
)

// SPI on the Feather M4.
var (
	SPI0 = SPI{
		Bus:    sam.SERCOM5_SPIM,
		SERCOM: 5,
	}

	SPI3 = SPI{
		Bus:    sam.SERCOM7_SPIM,
		SERCOM: 7,
	}
)

// USB CDC identifiers
const (
	usb_STRING_PRODUCT      = "Seeed Wio Terminal"
	usb_STRING_MANUFACTURER = "Seeed"
)

var (
	usb_VID uint16 = 0x2886
	usb_PID uint16 = 0x802D
)
