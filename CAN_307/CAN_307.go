package CAN_307

type CAN_307 struct {
	relayState byte
	func1      byte
	func2      byte
	func3      byte
	func4      byte
	func5      byte
	func6      byte
	func7      byte
}

func New(d []byte) CAN_307 {
	c := CAN_307{d[0], d[1], d[2], d[3], d[4], d[5], d[6], d[7]}
	return c
}

func (c CAN_307) Relay1_Master() bool {
	return (c.relayState & 1) != 0
}

func (c CAN_307) Relay2_Master() bool {
	return (c.relayState & 2) != 0
}

func (c CAN_307) Relay1_Slave1() bool {
	return (c.relayState & 4) != 0
}

func (c CAN_307) Relay2_Slave1() bool {
	return (c.relayState & 8) != 0
}

func (c CAN_307) Relay1_Slave2() bool {
	return (c.relayState & 0x10) != 0
}

func (c CAN_307) Relay2_Slave2() bool {
	return (c.relayState & 0x20) != 0
}

func (c CAN_307) GnRun() bool {
	return (c.func1 & 0x01) != 0
}

func (c CAN_307) GnRunSlave1() bool {
	return (c.func1 & 0x02) != 0
}

func (c CAN_307) GnRunSlave2() bool {
	return (c.func1 & 0x04) != 0
}

func (c CAN_307) AutoGn() bool {
	return (c.func2 & 0x01) != 0
}

func (c CAN_307) AutoLodExt() bool {
	return (c.func2 & 0x02) != 0
}

func (c CAN_307) AutoLodSoc() bool {
	return (c.func2 & 0x04) != 0
}

func (c CAN_307) Tm1() bool {
	return (c.func2 & 0x08) != 0
}

func (c CAN_307) Tm2() bool {
	return (c.func2 & 0x10) != 0
}

func (c CAN_307) ExtPwrDer() bool {
	return (c.func2 & 0x20) != 0
}

func (c CAN_307) ExtVfOk() bool {
	return (c.func2 & 0x40) != 0
}

func (c CAN_307) GdOn() bool {
	return (c.func2 & 0x80) != 0
}

func (c CAN_307) Error() bool {
	return (c.func3 & 0x01) != 0
}

func (c CAN_307) Run() bool {
	return (c.func3 & 0x02) != 0
}

func (c CAN_307) BatFan() bool {
	return (c.func3 & 0x04) != 0
}

func (c CAN_307) AcdCir() bool {
	return (c.func3 & 0x08) != 0
}

func (c CAN_307) MccBatFan() bool {
	return (c.func3 & 0x10) != 0
}

func (c CAN_307) MccAutoLod() bool {
	return (c.func3 & 0x20) != 0
}

func (c CAN_307) Chp() bool {
	return (c.func3 & 0x40) != 0
}

func (c CAN_307) ChpAdd() bool {
	return (c.func3 & 0x80) != 0
}

func (c CAN_307) SiComRemote() bool {
	return (c.func4 & 0x01) != 0
}

func (c CAN_307) Overload() bool {
	return (c.func4 & 0x02) != 0
}

func (c CAN_307) ExtSrcConn() bool {
	return (c.func5 & 0x01) != 0
}

func (c CAN_307) Silent() bool {
	return (c.func5 & 0x02) != 0
}

func (c CAN_307) Current() bool {
	return (c.func5 & 0x04) != 0
}

func (c CAN_307) FeedSelfC() bool {
	return (c.func5 & 0x08) != 0
}

func (c CAN_307) Esave() bool {
	return (c.func5 & 0x10) != 0
}
