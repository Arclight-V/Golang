package dbreader

import (
	"fmt"
	"os"
)

type DBreader interface {
	Read(file string) (*Recipes, error)
}

type JSONReader struct{}

func NewJSONReader() *JSONReader {
	return &JSONReader{}
}

func (j *JSONReader) Read(file string) (*Recipes, error) {
	f, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	fmt.Print(f)
	return &Recipes{}, nil
}

type XMLReader struct{}

func NewXMLReader() *XMLReader {
	return &XMLReader{}
}

func (x *XMLReader) Read(file string) (*Recipes, error) {
	f, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	fmt.Print(f)
	return &Recipes{}, nil
}
