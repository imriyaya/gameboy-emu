package cpu

type Registers struct {
	A  uint8
	B  uint8
	C  uint8
	D  uint8
	E  uint8
	F  uint8
	H  uint8
	L  uint8
	SP uint16
	PC uint16
}

func (r *Registers) GetAF() uint16 {
	return uint16(r.A)<<8 | uint16(r.F)
}

func (r *Registers) SetAF(value uint16) {
	r.A = uint8((value & 0xFF00) >> 8)
	r.F = uint8(value & 0xFF)
}

func (r *Registers) GetBC() uint16 {
	return uint16(r.B)<<8 | uint16(r.C)
}

func (r *Registers) SetBC(value uint16) {
	r.B = uint8((value & 0xFF00) >> 8)
	r.C = uint8(value & 0xFF)
}

func (r *Registers) GetDE() uint16 {
	return uint16(r.D)<<8 | uint16(r.E)
}

func (r *Registers) SetDE(value uint16) {
	r.D = uint8((value & 0xFF00) >> 8)
	r.E = uint8(value & 0xFF)
}

func (r *Registers) GetHL() uint16 {
	return uint16(r.H)<<8 | uint16(r.L)
}

func (r *Registers) SetHL(value uint16) {
	r.H = uint8((value & 0xFF00) >> 8)
	r.L = uint8(value & 0xFF)
}

const ZeroFlagBytePosition uint8 = 7
const SubtractFlagBytePosition uint8 = 6
const HalfCarryFlagBytePosition uint8 = 5
const CarryFlagBytePosition uint8 = 4

func (r *Registers) GetZero() bool {
	return ((r.F >> ZeroFlagBytePosition) & 0b1) != 0
}

func (r *Registers) SetZero(value bool) {
	if value {
		r.F |= 1 << ZeroFlagBytePosition
	} else {
		r.F &^= 1 << ZeroFlagBytePosition
	}
}

func (r *Registers) GetSubtract() bool {
	return ((r.F >> SubtractFlagBytePosition) & 0b1) != 0
}

func (r *Registers) SetSubtract(value bool) {
	if value {
		r.F |= 1 << SubtractFlagBytePosition
	} else {
		r.F &^= 1 << SubtractFlagBytePosition
	}
}

func (r *Registers) GetHalfCarry() bool {
	return ((r.F >> HalfCarryFlagBytePosition) & 0b1) != 0
}

func (r *Registers) SetHalfCarry(value bool) {
	if value {
		r.F |= 1 << HalfCarryFlagBytePosition
	} else {
		r.F &^= 1 << HalfCarryFlagBytePosition
	}
}

func (r *Registers) GetCarry() bool {
	return ((r.F >> CarryFlagBytePosition) & 0b1) != 0
}

func (r *Registers) SetCarry(value bool) {
	if value {
		r.F |= 1 << CarryFlagBytePosition
	} else {
		r.F &^= 1 << CarryFlagBytePosition
	}
}
