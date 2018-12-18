package chip8

type RAM []byte

func (r RAM) Write(pos int16, data []byte) error {
	for i, d := range data {
		r[pos+int16(i)] = d
	}
	return nil
}

func (r RAM) Read(pos uint16, size uint16) ([]byte, error) {
	return r[pos:pos+size], nil
}

func NewRAM() RAM{
	return make(RAM, 64)
}