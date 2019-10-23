package CAN_010

import "math"

type CAN_010 struct {
	fDecimal uint8
	fInteger uint8
}

func New(d []byte) CAN_010 {
	c := CAN_010{d[2], d[3]}
	return c
}

func (c CAN_010) Frequency() float64 {
	return math.Round(((float64(c.fDecimal)/256)+float64(c.fInteger))*100) / 100
}
