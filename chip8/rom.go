package chip8

import (
	"errors"
)

type ROM []byte

func (r ROM) Write(pos int16, data []byte) error {
	return errors.New("Unsupported operation. You can't write to a ROM")
}

func (r ROM) Read(pos uint16, size uint16) ([]byte, error) {
	if int(pos) >= len(r) || int(pos+size) >= len(r) {
		return nil, errors.New("Memory position not found")
	}
	return r[pos : pos+size], nil
}

func NewROM() ROM {
	return make(ROM, 64)
}
