package kml

import (
	"encoding/xml"
	"github.com/vc2402/kml-reader/models"
	"io"
)

// Read reads kml file from io.Reader and returns an unmarshalled document
func Read(source io.Reader) (models.KML, error) {
	xmlDecoder := xml.NewDecoder(source)
	var kml models.KML
	err := xmlDecoder.Decode(&kml)
	if err != nil {
		return models.KML{}, err
	}
	return kml, nil
}
