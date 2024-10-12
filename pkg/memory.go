package pkg

type Memory struct {
	memory [0xFFFF]uint8
}

func New() *Memory {
	mem := &Memory{}
	return mem
}

func (mem *Memory) GetValue(address uint16) uint8 {
	return mem.memory[address]
}

func (mem *Memory) SetValue(address uint16, value uint8) {
	mem.memory[address] = value
}

func (mem *Memory) GetPointer(address uint16) *uint8 {
	return &mem.memory[address]
}
