package midi

import (
	"machine"
)

const (
	midiEndpointOut = 5 // from PC
	midiEndpointIn  = 6 // to PC
)

var Midi *midi

type midi struct {
	buf *RingBuffer
}

func init() {
	if Midi == nil {
		Midi = newMidi()
		machine.EnableMIDI(Midi.Callback)
	}
}

// New returns hid-mouse.
func New() *midi {
	return Midi
}

func newMidi() *midi {
	return &midi{
		buf: NewRingBuffer(),
	}
}

func (m *midi) SetCallback(callbackRx func([]byte)) {
	machine.SetCallbackMIDI(callbackRx)
}

// SendUSBPacket sends a HIDPacket.
func (m *midi) SendUSBPacket(b []byte) {
	machine.SendUSBMIDIPacket(midiEndpointIn, b)
}

func (m *midi) Callback() {
	if b, ok := m.buf.Get(); ok {
		m.SendUSBPacket(b)
	}
}
