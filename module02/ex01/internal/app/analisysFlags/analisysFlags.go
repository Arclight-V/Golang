package analisysflags

import "fmt"

type FlagMode uint32

// define a flag
const (
	L FlagMode = iota
	M
	W
)

type Flags struct {
	flag          FlagMode
	analisisFiles []string
}

func NewFlags(f FlagMode, aFiles []string) *Flags {
	return &Flags{flag: f, analisisFiles: aFiles}
}

func (f *Flags) Mode() FlagMode {
	return f.flag
}

func (f *Flags) Files() []string {
	return f.analisisFiles
}

func (f *Flags) LenFiles() int {
	return len(f.analisisFiles)
}

func (f *Flags) String() string {
	var mode rune
	switch {
	case f.flag == L:
		mode = 'L'
	case f.flag == M:
		mode = 'M'
	case f.flag == W:
		mode = 'W'
	default:
		mode = 'N'
	}
	return fmt.Sprintf("mode: %c\t files: %v", mode, f.analisisFiles)
}
