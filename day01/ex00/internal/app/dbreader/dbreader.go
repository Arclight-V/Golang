package dbreader

type DBreader interface {
	Read(file string)
}

type XMLReader struct {
	reader DBreader
}

func NewXMLreader() *XMLReader {
	return &XMLReader{}
}

type JSONReader struct {
	reader DBreader
}
