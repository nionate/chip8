package chip8

import "fmt"

const InstructionSize = 2
const MemoryStart = 0x200
const MemorySize = 16

type CPU struct {
	V     [MemorySize]uint8
	I     uint16
	PC    uint16
	SP    uint8
	Stack [MemorySize]uint16
	//TO-DO: implements timmers
}

func (cpu *CPU) Run(cart *Cartridge) error {
	for true {
		data, err := cart.Read(cpu.PC, InstructionSize)
		if err != nil {
			return err
		}
		cpu.Interpret(data)
		cpu.PC += InstructionSize
	}
	return nil
}

func (cpu *CPU) Interpret(data []byte) error {
	instr := uint16(data[0])<<8 | uint16(data[1])
	//fmt.Printf("PC = %04X -> %04X\n", cpu.PC, instr)

	switch instr >> 12 & 0xF { // check the first 4 bits
	case 1:
	case 2:
	case 3:
		return SE(cpu, instr)
	case 4:
		return SNE(cpu, instr)
	case 5:
		return SE5(cpu, instr)
	case 6:
		return LD(cpu, instr)
	case 7:
	case 8:
		switch instr & 0xF { // check the last 4 bits
		case 0:
			fmt.Printf("%04X => Set Vx = Vy.\n", instr)
		case 1:
			fmt.Printf("%04X => Set Vx = Vx OR Vy.\n", instr)
		}
	case 0xF:
		switch instr & 0xFF { // check the last byte
		case 0x0A:
			fmt.Printf("%04X => Wait for a key press, store the value of the key in Vx.\n", instr)
		case 0x29:
			fmt.Printf("%04X => Set I = location of sprite for digit Vx.\n", instr)
		}
	}
	return nil
}

func NewCPU() *CPU {
	return &CPU{
		PC: MemoryStart, //program start here
	}
}
