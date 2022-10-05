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
	Flag          FlagMode
	AnalisisFiles []string
}

func (f *Flags) String() string {
	var mode rune
	switch {
	case f.Flag == L:
		mode = 'L'
	case f.Flag == M:
		mode = 'M'
	case f.Flag == W:
		mode = 'W'
	default:
		mode = 'N'
	}
	return fmt.Sprintf("mode: %c\t files: %v", mode, f.AnalisisFiles)
}
