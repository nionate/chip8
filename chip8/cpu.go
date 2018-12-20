package chip8

import "fmt"

const InstructionSize = 2
const MemoryStart = 0x200

type CPU struct {
	V     [0xF]uint8
	I     uint16
	PC    uint16
	SP    uint8
	Stack [0xF]uint16
	//TO-DO: implements timmers
}

func (cpu *CPU) Run(cart *Cartridge) error {
	for true {
		instruction, err := cart.Read(cpu.PC, InstructionSize)
		if err != nil {
			return err
		}
		fmt.Printf("%X\n", instruction)

		cpu.PC += InstructionSize
	}
	return nil
}

func NewCPU() *CPU {
	return &CPU{
		PC: MemoryStart, //program start here
	}
}
