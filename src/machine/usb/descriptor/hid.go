package descriptor

import (
	"bytes"
	"encoding/binary"
	"errors"
)

var configurationCDCHID = [configurationTypeLen]byte{
	configurationTypeLen,
	TypeConfiguration,
	0x84, 0x00, // adjust length as needed
	0x04, // number of interfaces
	0x01, // configuration value
	0x00, // index to string description
	0xa0, // attributes
	0x32, // maxpower
}

var ConfigurationCDCHID = ConfigurationType{
	data: configurationCDCHID[:],
}

var interfaceHID = [interfaceTypeLen]byte{
	interfaceTypeLen,
	TypeInterface,
	0x02, // InterfaceNumber
	0x00, // AlternateSetting
	0x01, // NumEndpoints
	0x03, // InterfaceClass
	0x00, // InterfaceSubClass
	0x00, // InterfaceProtocol
	0x00, // Interface
}

var InterfaceHID = InterfaceType{
	data: interfaceHID[:],
}

const (
	ClassHIDTypeLen = 9
)

type ClassHIDType struct {
	data []byte
}

func (d ClassHIDType) Bytes() []byte {
	return d.data[:]
}

func (d ClassHIDType) Length(v uint8) {
	d.data[0] = byte(v)
}

func (d ClassHIDType) Type(v uint8) {
	d.data[1] = byte(v)
}

func (d ClassHIDType) HID(v uint16) {
	binary.LittleEndian.PutUint16(d.data[2:4], v)
}

func (d ClassHIDType) CountryCode(v uint8) {
	d.data[4] = byte(v)
}

func (d ClassHIDType) NumDescriptors(v uint8) {
	d.data[5] = byte(v)
}

func (d ClassHIDType) ClassType(v uint8) {
	d.data[6] = byte(v)
}

func (d ClassHIDType) ClassLength(v uint16) {
	binary.LittleEndian.PutUint16(d.data[7:9], v)
}

var errNoClassHIDFound = errors.New("no classHID found")

// FindClassHIDType tries to find the ClassHID class in the descriptor.
func FindClassHIDType(des, class []byte) (ClassHIDType, error) {
	if len(des) < ClassHIDTypeLen || len(class) == 0 {
		return ClassHIDType{}, errNoClassHIDFound
	}

	// search only for ClassHIDType without the ClassLength,
	// in case it has already been set.
	idx := bytes.Index(des, class[:ClassHIDTypeLen-2])
	if idx == -1 {
		return ClassHIDType{}, errNoClassHIDFound
	}

	return ClassHIDType{data: des[idx : idx+ClassHIDTypeLen]}, nil
}

var classHID = [ClassHIDTypeLen]byte{
	ClassHIDTypeLen,
	TypeClassHID,
	0x11, // HID version L
	0x01, // HID version H
	0x00, // CountryCode
	0x01, // NumDescriptors
	0x22, // ClassType
	0x7E, // ClassLength L
	0x00, // ClassLength H
}

var ClassHID = ClassHIDType{
	data: classHID[:],
}

var CDCHID = Descriptor{
	Device: DeviceCDC.Bytes(),
	Configuration: Append([][]byte{
		ConfigurationCDCHID.Bytes(),
		InterfaceAssociationCDC.Bytes(),
		InterfaceCDCControl.Bytes(),
		ClassSpecificCDCHeader.Bytes(),
		ClassSpecificCDCACM.Bytes(),
		ClassSpecificCDCUnion.Bytes(),
		ClassSpecificCDCCallManagement.Bytes(),
		EndpointEP1IN.Bytes(),
		InterfaceCDCData.Bytes(),
		EndpointEP2OUT.Bytes(),
		EndpointEP3IN.Bytes(),
		InterfaceHID.Bytes(),
		ClassHID.Bytes(),
		EndpointEP4IN.Bytes(),
		[]byte{
			// 32 byte
			0x09, 0x04, 0x03, 0x00, 0x02, 0x03, 0x00, 0x00, 0x00,
			0x09, 0x21, 0x11, 0x01, 0x00, 0x01, 0x22, 0x22, 0x00,
			0x07, 0x05, 0x86, 0x03, 0x20, 0x00, 0x01,
			0x07, 0x05, 0x07, 0x03, 0x20, 0x00, 0x01,
		},
	}),
	HID: map[uint16][]byte{
		2: Append([][]byte{
			HIDUsagePageGenericDesktop,
			HIDUsageDesktopKeyboard,
			HIDCollectionApplication,
			HIDReportID(2),

			HIDUsagePageKeyboard,
			HIDUsageMinimum(224),
			HIDUsageMaximum(231),
			HIDLogicalMinimum(0),
			HIDLogicalMaximum(1),
			HIDReportSize(1),
			HIDReportCount(8),
			HIDInputDataVarAbs,
			HIDReportCount(1),
			HIDReportSize(8),
			HIDInputConstVarAbs,
			HIDReportCount(6),
			HIDReportSize(8),
			HIDLogicalMinimum(0),
			HIDLogicalMaximum(255),

			HIDUsagePageKeyboard,
			HIDUsageMinimum(0),
			HIDUsageMaximum(255),
			HIDInputDataAryAbs,
			HIDCollectionEnd,

			HIDUsagePageGenericDesktop,
			HIDUsageDesktopMouse,
			HIDCollectionApplication,
			HIDUsageDesktopPointer,
			HIDCollectionPhysical,
			HIDReportID(1),

			HIDUsagePageButton,
			HIDUsageMinimum(1),
			HIDUsageMaximum(5),
			HIDLogicalMinimum(0),
			HIDLogicalMaximum(1),
			HIDReportCount(5),
			HIDReportSize(1),
			HIDInputDataVarAbs,
			HIDReportCount(1),
			HIDReportSize(3),
			HIDInputConstVarAbs,

			HIDUsagePageGenericDesktop,
			HIDUsageDesktopX,
			HIDUsageDesktopY,
			HIDUsageDesktopWheel,
			HIDLogicalMinimum(-127),
			HIDLogicalMaximum(127),
			HIDReportSize(8),
			HIDReportCount(3),
			HIDInputDataVarRel,
			HIDCollectionEnd,
			HIDCollectionEnd,

			HIDUsagePageConsumer,
			HIDUsageConsumerControl,
			HIDCollectionApplication,
			HIDReportID(3),
			HIDLogicalMinimum(0),
			HIDLogicalMaximum(8191),
			HIDUsageMinimum(0),
			HIDUsageMaximum(0x1FFF),
			HIDReportSize(16),
			HIDReportCount(1),
			HIDInputDataAryAbs,
			HIDCollectionEnd}),
		3: []byte{
			0x06, 0x60, 0xff, // Usage Page (Vendor-Defined 0xFF60)
			0x09, 0x61, // Usage (Vendor-Defined 0x61)
			0xa1, 0x01, // Collection (Application)
			0x09, 0x62, //   Usage (Vendor-Defined 0x62)
			0x15, 0x00, //   Logical Minimum (0)
			0x26, 0xff, 0x00, //   Logical Maximum (255)
			0x95, 0x20, //   Report Count (32)
			0x75, 0x08, //   Report Size (8)
			0x81, 0x02, //   Input (Data, Var, Abs)
			0x09, 0x63, //   Usage (Vendor-Defined 0x63)
			0x15, 0x00, //   Logical Minimum (0)
			0x26, 0xff, 0x00, //   Logical Maximum (255)
			0x95, 0x20, //   Report Count (32)
			0x75, 0x08, //   Report Size (8)
			0x91, 0x02, //   Output (Data, Var, Abs)
			0xc0, // End Collection
		},
	},
}
