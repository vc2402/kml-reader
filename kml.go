package kml

import (
	"io"

	"github.com/zuchi/kml-go/v1/internal"
	"github.com/zuchi/kml-go/v1/models/kml"
)

type KmlReader interface {
	GetOuterPolygon() ([]kml.PolygonData, error)
}

func NewKmlManager(kmlBody io.Reader) (KmlReader, error) {
	return internal.NewKmlManager(kmlBody)
}
