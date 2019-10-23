package CAN_305

import "encoding/binary"

type CAN_305 struct {
	vBatt   uint16
	iBatt   int16
	tBatt   int16
	socBatt uint16
}

func New(d []byte) CAN_305 {
	c := CAN_305{binary.LittleEndian.Uint16(d[0:]), int16(binary.LittleEndian.Uint16(d[2:])), int16(binary.LittleEndian.Uint16(d[4:])), binary.LittleEndian.Uint16(d[6:])}
	return c
}

func (c CAN_305) VBatt() float32 {
	return float32(c.vBatt) / 10
}

func (c CAN_305) IBatt() float32 {
	return float32(c.iBatt) / 10
}

func (c CAN_305) TBatt() float32 {
	return float32(c.tBatt) / 10
}

func (c CAN_305) SocBatt() float32 {
	return float32(c.socBatt) / 10
}
