package tests

import (
	"ex00/service/dbreader"
	"testing"
)

func Process(obj dbreader.DBreader, path string) (*dbreader.Recipes, error) {
	return obj.Read(path)
}

func TestDBReaderImpl(t *testing.T) {
	pathToXML := "testDB/test_original_database.xml"
	xmlReader := dbreader.NewXMLReader()

	_, err := Process(xmlReader, pathToXML)
	if err != nil {
		t.Errorf("Incorect result. Expected %#v, you result %s\"", nil, err)
	}

	pathToJSON := "testDB/test_stolen_database.json"
	JSONReader := dbreader.NewJSONReader()

	_, e := Process(JSONReader, pathToJSON)
	if e != nil {
		t.Errorf("Incorect result. Expected %#v, you result %s\"", nil, e)
	}

}
