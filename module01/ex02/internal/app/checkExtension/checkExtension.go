package checkextension

import (
	"path/filepath"
)

func Checkextension(fileToPath string, extension string) bool {

	return filepath.Ext(fileToPath) == extension

}
