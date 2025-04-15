package utils

import (
	"github.com/lestrrat-go/libxml2"
	"github.com/lestrrat-go/libxml2/xsd"
)

func ValidateXML(xmlData []byte, schemaPath string) error {
	schema, err := xsd.ParseFromFile(schemaPath)
	if err != nil {
		return err
	}
	defer schema.Free()

	doc, err := libxml2.Parse(xmlData)
	if err != nil {
		return err
	}
	defer doc.Free()

	return schema.Validate(doc)
}
