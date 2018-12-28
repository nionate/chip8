package main

import (
	"fmt"
	"log"

	"github.com/nionate/chip8/chip8"
)

func main() {
	cartridge, err := chip8.LoadCartridge("./roms/pong2.rom")
	if err != nil {
		log.Panic(err)
	}
	cpu := chip8.CPU{}
	err = cpu.Run(cartridge)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Chip8")
}
