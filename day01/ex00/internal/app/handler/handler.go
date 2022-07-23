package handler

import "ex00/service/dbreader"

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

// TODO add find extension file method
func (h *Handler) Handle(file string) (*dbreader.Recipes, error) {
	switch file {
	case ".xml":
		h.xmlReader.Read(file)
	case ".json":
		h.jsonReader.Read(file)
	default:

	}
}
