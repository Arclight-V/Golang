package handler

import (
	"dbreader"
	"errors"
	"strings"
)

var extensions = [...]string{
	".xml",
	".json",
}

type Handler struct {
	jsonReader *dbreader.JSONReader
	xmlReader  *dbreader.XMLReader
}

func NewHandler() *Handler {
	return &Handler{
		jsonReader: dbreader.NewJSONReader(),
		xmlReader:  dbreader.NewXMLReader(),
	}
}

func FindExtension(path string) string {
	for _, e := range extensions {
		if is_find := strings.Contains(path, e); is_find == true {
			return e
		}
	}
	return "none"
}

func (h *Handler) Handle(file string) (*dbreader.Recipes, error) {
	switch FindExtension(file) {
	case ".xml":
		return h.xmlReader.Read(file)
	case ".json":
		return h.jsonReader.Read(file)
	default:
	}
	return nil, errors.New("Error! Unreachable file extension")
}
