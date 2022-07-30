package dbreader

import (
	"encoding/json"
	"encoding/xml"
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
	recipes := Recipes{}
	if e := json.Unmarshal(f, &recipes); e != nil {
		return nil, e
	}

	return &recipes, nil
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
	recipes := Recipes{}
	if e := xml.Unmarshal(f, &recipes); e != nil {
		return nil, e
	}
	return &recipes, nil
}
