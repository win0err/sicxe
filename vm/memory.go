package vm

const MemorySize = 1 << 15        // (1 << 20) in SIX/XE
const MaxAddress = MemorySize - 1 // SIC

const (
	BootstrapAddress int = 0x00
	LoaderAddress    int = 0x80
	ShellAddress     int = 0x300
	UserAddress      int = 0x1000
	SystemAddress    int = 0x00
)

type Memory = []Byte

func NewMemory() Memory {
	return make(Memory, MaxAddress)
}
