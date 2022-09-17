package flagsSerach

type Flags struct {
	IsSymlinks, IsDirectories, IsFile bool
	FileExtension, Folder             string
}

func (f *Flags) AllFlagFalse() bool {
	if !f.IsSymlinks && !f.IsDirectories && !f.IsFile {
		return true
	}
	return false
}

func (f *Flags) OnlyDir() bool {
	if !f.IsSymlinks && f.IsDirectories && !f.IsFile {
		return trues
	}
	return false
}
