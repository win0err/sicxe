package vm

type Registers struct {
	A  Word // Accumulator
	X  Word // Index Register
	L  Word // Linkage Register
	PC Word // Program Counter
	SW Word // Status Word
}

func NewRegisters() *Registers {
	return &Registers{
		A:  0xFFFFFF,
		X:  0xFFFFFF,
		L:  0xFFFFFF,
		PC: 0x000000,
		SW: 0xFFFF3F,
	}
}

const MaxDevices = 7 // 255 in SIX/XE

// Device types
const (
	BootDevice   int = 0
	InputDevice  int = 1
	OutputDevice int = 2
)
