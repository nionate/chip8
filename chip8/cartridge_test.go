package chip8_test

import (
	"bytes"
	"testing"

	"github.com/nionate/chip8/chip8"
)

func TestShouldLoadACartridge(t *testing.T) {
	filename := "../roms/pong2.rom"
	cart, err := chip8.LoadCartridge(filename)
	if err != nil {
		t.Error(err)
	}

	if cart.Filename != filename {
		t.Error("Invalid filename")
	}

	if len(cart.ROM) == 0 {
		t.Error("Unexpected cartridge size")
	}
}

func TestShouldNotWriteInCartridge(t *testing.T) {
	filename := "../roms/pong2.rom"
	cart, err := chip8.LoadCartridge(filename)
	if err != nil {
		t.Error(err)
	}

	errors := cart.Write(4, []byte{1})
	if errors == nil {
		t.Error(err)
	}
}

func TestShouldReadFromCartridge(t *testing.T) {
	filename := "../roms/pong2.rom"
	cart, err := chip8.LoadCartridge(filename)
	if err != nil {
		t.Error(err)
	}

	instruction, err := cart.Read(4, 2)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(instruction, []byte{0x6C, 0x3F}) {
		t.Errorf("Unexpected bytes: %X", instruction)
	}
}

func TestShouldNotReadFromCartridgeWhenRomDoesNotExists(t *testing.T) {
	filename := "../roms/pong22.rom"
	_, err := chip8.LoadCartridge(filename)
	if err == nil {
		t.Error(err)
	}
}
