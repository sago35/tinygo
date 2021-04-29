// +build sam,atsame51 sam,atsame54

package machine

import (
	"device/sam"
	"errors"
	"unsafe"
)

// Message RAM can only be located in the first 64 KB area of the system RAM.

//go:align 4
var CANRxFifo [2][(8 + 64) * 16]byte

//go:align 4
var CANTxFifo [2][(8 + 64) * 16]byte

//go:align 4
var CANEvFifo [2][(8 + 64) * 16]byte

type CAN struct {
	Bus     *sam.CAN_Type
	Channel byte
}

type CANTransferRate uint32

const (
	CANTransferRate125kbps  CANTransferRate = 125000
	CANTransferRate250kbps  CANTransferRate = 250000
	CANTransferRate500kbps  CANTransferRate = 500000
	CANTransferRate1000kbps CANTransferRate = 1000000
	CANTransferRate2000kbps CANTransferRate = 2000000
	CANTransferRate4000kbps CANTransferRate = 4000000
)

type CANConfig struct {
	TransferRate   CANTransferRate
	TransferRateFD CANTransferRate
	Tx             Pin
	Rx             Pin
}

var (
	errCANInvalidTransferRate   = errors.New("CAN: invalid TransferRate")
	errCANInvalidTransferRateFD = errors.New("CAN: invalid TransferRateFD")
)

func (can *CAN) Configure(config CANConfig) error {
	can.Bus.CCCR.SetBits(sam.CAN_CCCR_INIT)
	for !can.Bus.CCCR.HasBits(sam.CAN_CCCR_INIT) {
	}

	can.Bus.CCCR.SetBits(sam.CAN_CCCR_CCE)
	can.Channel = 0
	if can.Bus == sam.CAN1 {
		can.Channel = 1
	}

	can.Bus.CCCR.SetBits(sam.CAN_CCCR_BRSE | sam.CAN_CCCR_FDOE)
	can.Bus.MRCFG.Set(sam.CAN_MRCFG_QOS_MEDIUM)
	// base clock == 48 MHz
	if config.TransferRate == 0 {
		config.TransferRate = CANTransferRate500kbps
	}
	brp := uint32(6)
	switch config.TransferRate {
	case CANTransferRate125kbps:
		brp = 32
	case CANTransferRate250kbps:
		brp = 16
	case CANTransferRate500kbps:
		brp = 8
	case CANTransferRate1000kbps:
		brp = 4
	default:
		return errCANInvalidTransferRate
	}
	can.Bus.NBTP.Set(8<<sam.CAN_NBTP_NTSEG1_Pos | (brp-1)<<sam.CAN_NBTP_NBRP_Pos |
		1<<sam.CAN_NBTP_NTSEG2_Pos | 3<<sam.CAN_NBTP_NSJW_Pos)

	if config.TransferRateFD == 0 {
		config.TransferRateFD = CANTransferRate1000kbps
	}
	if config.TransferRateFD < config.TransferRate {
		return errCANInvalidTransferRateFD
	}

	brp = uint32(2)
	switch config.TransferRateFD {
	case CANTransferRate125kbps:
		brp = 32
	case CANTransferRate250kbps:
		brp = 16
	case CANTransferRate500kbps:
		brp = 8
	case CANTransferRate1000kbps:
		brp = 4
	case CANTransferRate2000kbps:
		brp = 2
	case CANTransferRate4000kbps:
		brp = 1
	default:
		return errCANInvalidTransferRateFD
	}
	can.Bus.DBTP.Set((brp-1)<<sam.CAN_DBTP_DBRP_Pos | 8<<sam.CAN_DBTP_DTSEG1_Pos |
		1<<sam.CAN_DBTP_DTSEG2_Pos | 3<<sam.CAN_DBTP_DSJW_Pos)

	switch config.TransferRateFD {
	case CANTransferRate1000kbps:
	default:
	}

	can.Bus.RXF0C.Set(sam.CAN_RXF0C_F0OM | 16<<sam.CAN_RXF0C_F0S_Pos | uint32(uintptr(unsafe.Pointer(&CANRxFifo[can.Channel][0])))&0xFFFF)
	can.Bus.RXESC.Set(sam.CAN_RXESC_F0DS_DATA64)
	can.Bus.TXESC.Set(sam.CAN_TXESC_TBDS_DATA64)
	can.Bus.TXBC.Set(16<<sam.CAN_TXBC_TFQS_Pos | 0<<sam.CAN_TXBC_NDTB_Pos | uint32(uintptr(unsafe.Pointer(&CANTxFifo[can.Channel][0])))&0xFFFF)
	can.Bus.TXEFC.Set(2<<sam.CAN_TXEFC_EFS_Pos | uint32(uintptr(unsafe.Pointer(&CANEvFifo[can.Channel][0])))&0xFFFF)

	can.Bus.TSCC.Set(sam.CAN_TSCC_TSS_INC)

	can.Bus.GFC.Set(0<<sam.CAN_GFC_ANFS_Pos | 0<<sam.CAN_GFC_ANFE_Pos)

	can.Bus.SIDFC.Set(0 << sam.CAN_SIDFC_LSS_Pos)
	can.Bus.XIDFC.Set(0 << sam.CAN_SIDFC_LSS_Pos)

	can.Bus.XIDAM.Set(0x1FFFFFFF << sam.CAN_XIDAM_EIDM_Pos)

	can.Bus.CCCR.ClearBits(sam.CAN_CCCR_CCE)
	can.Bus.CCCR.ClearBits(sam.CAN_CCCR_INIT)
	for can.Bus.CCCR.HasBits(sam.CAN_CCCR_INIT) {
	}

	mode := PinCAN0
	if can.Bus == sam.CAN1 {
		mode = PinCAN1
	}
	config.Rx.Configure(PinConfig{Mode: mode})
	config.Tx.Configure(PinConfig{Mode: mode})

	return nil
}

func (can CAN) Tx(id uint32, dlc byte, data []byte) {
	putIndex := (can.Bus.TXFQS.Get() & sam.CAN_TXFQS_TFQPI_Msk) >> sam.CAN_TXFQS_TFQPI_Pos

	f := CANTxFifo[can.Channel][putIndex*(8+64) : (putIndex+1)*(8+64)]
	if true {
		// standard identifier is stored into ID[28:18]
		id <<= 18
	}

	for i := byte(0); i < dlc; i++ {
		f[8+i] = data[i]
	}

	f[3] = byte(id >> 24)
	f[2] = byte(id >> 16)
	f[1] = byte(id >> 8)
	f[0] = byte(0)
	f[7] = 0x00
	f[6] = dlc
	f[5] = 0x00 // reserved
	f[4] = 0x00 // reserved
	can.Bus.TXBAR.SetBits(1 << putIndex)
}

func (can CAN) TxFifoIsFull() bool {
	return (can.Bus.TXFQS.Get() & sam.CAN_TXFQS_TFQF_Msk) == sam.CAN_TXFQS_TFQF_Msk
}

func (can CAN) TxFD(id uint32, dlc byte, data []byte) {
	putIndex := (can.Bus.TXFQS.Get() & sam.CAN_TXFQS_TFQPI_Msk) >> sam.CAN_TXFQS_TFQPI_Pos

	f := CANTxFifo[can.Channel][putIndex*(8+64) : (putIndex+1)*(8+64)]
	if true {
		// standard identifier is stored into ID[28:18]
		id <<= 18
	}

	for i := byte(0); i < dlc; i++ {
		f[8+i] = data[i]
	}

	if dlc <= 0x08 {
	} else if dlc <= 12 {
		dlc = 0x09
	} else if dlc <= 16 {
		dlc = 0x0A
	} else if dlc <= 20 {
		dlc = 0x0B
	} else if dlc <= 24 {
		dlc = 0x0C
	} else if dlc <= 32 {
		dlc = 0x0D
	} else if dlc <= 48 {
		dlc = 0x0E
	} else if dlc <= 64 {
		dlc = 0x0F
	}

	f[3] = byte(id >> 24)
	f[2] = byte(id >> 16)
	f[1] = byte(id >> 8)
	f[0] = byte(0)
	f[7] = 0x00
	f[6] = dlc | 0x20 | 0x10
	f[5] = 0x00 // reserved
	f[4] = 0x00 // reserved
	can.Bus.TXBAR.SetBits(1 << putIndex)
}

func (can CAN) Size() int {
	sz := (can.Bus.RXF0S.Get() & sam.CAN_RXF0S_F0FL_Msk) >> sam.CAN_RXF0S_F0FL_Pos
	return int(sz)
}

func (can CAN) Rx() (id uint32, dlc byte, data []byte) {
	idx := (can.Bus.RXF0S.Get() & sam.CAN_RXF0S_F0GI_Msk) >> sam.CAN_RXF0S_F0GI_Pos
	f := CANRxFifo[can.Channel][idx*(8+64):]

	id = ((uint32(f[3]) << 24) + (uint32(f[2]) << 16) + (uint32(f[1]) << 8) + uint32(f[0])) & 0x1FFFFFFF
	if (f[3] & 0x20) == 0 {
		id >>= 18
		id &= 0x000007FF
	}
	dlc = byte(idx)

	dlc = f[6] & 0x0F
	if dlc == 0x09 {
		dlc = 12
	} else if dlc == 0x0A {
		dlc = 16
	} else if dlc == 0x0B {
		dlc = 20
	} else if dlc == 0x0C {
		dlc = 24
	} else if dlc == 0x0D {
		dlc = 32
	} else if dlc == 0x0E {
		dlc = 48
	} else if dlc == 0x0F {
		dlc = 64
	}

	for i := byte(0); i < dlc; i++ {
		data = append(data, f[i+8])
	}
	data = data[:dlc]

	can.Bus.RXF0A.ReplaceBits(idx, sam.CAN_RXF0A_F0AI_Msk, sam.CAN_RXF0A_F0AI_Pos)

	return id, dlc, data
}
