package internal

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/zuchi/kml-go/v1/models"
	"github.com/zuchi/kml-go/v1/models/kml"
)

var (
	ErrKmlManagerIsNil      = errors.New("there is no kml manager object")
	ErrOuterPolygonNotFound = errors.New("there is no outer polygon in the provided kml")
	ErrLatConversion        = errors.New("latitude isn't valid.")
	ErrLonConversion        = errors.New("longitude isn't valid.")
	ErrAltConversion        = errors.New("altitude isn't valid.")
)

type KmlManager struct {
	kmlDoc models.KML
}

func NewKmlManager(kmlBody io.Reader) (*KmlManager, error) {
	body, err := io.ReadAll(kmlBody)
	if err != nil {
		return nil, err
	}

	var doc models.KML
	if err := xml.Unmarshal(body, &doc); err != nil {
		return nil, fmt.Errorf("couldn't parse kml file into KML model: %w", err)
	}

	k := &KmlManager{
		kmlDoc: doc,
	}

	return k, nil
}

func (km *KmlManager) GetOuterPolygon() ([]kml.PolygonData, error) {
	areas := make(map[string][]kml.GeoPoint, 0)
	polygons := []kml.PolygonData{}
	if km == nil {
		return polygons, ErrKmlManagerIsNil
	}

	doc := km.kmlDoc.Document.Placemarks
	for _, p := range doc {
		if p.Polygon != nil && p.Polygon.OuterBoundaryIs != nil {
			lat, lon, alt, err := convertLatLonAltFromItem(p.Polygon.OuterBoundaryIs.Coordinates)
			if err != nil {
				return polygons, err
			}
			coordinates, found := areas[p.Name]
			if found {
				coordinates = append(coordinates, kml.GeoPoint{
					Latitude:  lat,
					Longitude: lon,
					Altitude:  alt,
				})

				areas[p.Name] = coordinates
				continue
			}

			areas[p.Name] = []kml.GeoPoint{
				{
					Latitude:  lat,
					Longitude: lon,
					Altitude:  alt,
				},
			}
		}
	}

	for k, v := range areas {
		polygons = append(polygons, kml.PolygonData{
			Name:        k,
			Coordinates: v,
		})
	}

	return polygons, nil
}

func convertLatLonAltFromItem(coordinates string) (float64, float64, float64, error) {
	item := strings.Split(coordinates, ",")
	lon, err := strconv.ParseFloat(item[0], 64)
	if err != nil {
		return 0, 0, 0, ErrLonConversion
	}
	lat, err := strconv.ParseFloat(item[1], 64)
	if err != nil {
		return 0, 0, 0, ErrLatConversion
	}
	alt, err := strconv.ParseFloat(item[2], 64)
	if err != nil {
		return 0, 0, 0, ErrAltConversion
	}

	return lat, lon, alt, nil
}
