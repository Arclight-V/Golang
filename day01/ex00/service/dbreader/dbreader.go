package dbreader

import "errors"

type DBreader interface {
	Read(file string) (*Recipes, error)
}

type JSONReader struct{}

func NewJSONReader() *JSONReader {
	return &JSONReader{}
}

func (j *JSONReader) Read(file string) (*Recipes, error) {
	if len(file) == 0 {
		return nil, errors.New("Error")
	}
	return &Recipes{}, nil
}

type XMLReader struct{}

func NewXMLReader() *XMLReader {
	return &XMLReader{}
}

func (x *XMLReader) Read(file string) (*Recipes, error) {
	if len(file) == 0 {
		return nil, errors.New("Error")
	}
	return &Recipes{}, nil
}
