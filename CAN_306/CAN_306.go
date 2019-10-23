package CAN_306

import "encoding/binary"

type CAN_306 struct {
	sohBatt            uint16
	chargingProcedure  byte
	operatingState     byte
	activeErrorMessage uint16
	chargeSetPoint     uint16
}

func New(d []byte) CAN_306 {
	c := CAN_306{binary.LittleEndian.Uint16(d[0:]), d[2], d[3], binary.LittleEndian.Uint16(d[4:]), binary.LittleEndian.Uint16(d[6:])}
	return c
}

func (c CAN_306) SohBatt() float32 {
	return (float32(c.sohBatt) / 10)
}

func (c CAN_306) ChargingProcedure() byte {
	return c.chargingProcedure
}

func (c CAN_306) OperatingState() byte {
	return c.operatingState
}

func (c CAN_306) ActiveErrorMessage() uint16 {
	return c.activeErrorMessage
}

func (c CAN_306) ChargeSetPoint() float32 {
	return (float32(c.chargeSetPoint) / 10)
}
