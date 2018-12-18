package chip8

import "fmt"

const InstructionSize = 2

type CPU struct {
	PC uint16
}

func (cpu CPU) Run(cart *Cartridge) error{
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