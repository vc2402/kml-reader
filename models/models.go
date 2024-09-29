package models

import (
	"encoding/xml"
	"time"
)

// KML represents the root element of a KML document.
type KML struct {
	XMLName xml.Name `xml:"kml"`
	Xmlns   string   `xml:"xmlns,attr"`

	Document  *Document  `xml:"Document,omitempty"`
	Placemark *Placemark `xml:"Placemark,omitempty"`
	// Add more elements as necessary
}

// Document represents a KML Document element.
type Document struct {
	Name        string      `xml:"name,omitempty"`
	Description string      `xml:"description,omitempty"`
	Placemarks  []Placemark `xml:"Placemark"`
	Folders     []Folder    `xml:"Folder,omitempty"`
	// Add more elements as necessary
}

// Folder represents a KML Folder element.
type Folder struct {
	Name        string      `xml:"name,omitempty"`
	Description string      `xml:"description,omitempty"`
	Placemarks  []Placemark `xml:"Placemark"`
	// Add more elements as necessary
}

// Placemark represents a KML Placemark element.
type Placemark struct {
	Name        string      `xml:"name,omitempty"`
	Description string      `xml:"description,omitempty"`
	Point       *Point      `xml:"Point,omitempty"`
	LineString  *LineString `xml:"LineString,omitempty"`
	Polygon     *Polygon    `xml:"Polygon,omitempty"`
	// Add more geometry types as necessary
}

// Point represents a KML Point element.
type Point struct {
	Coordinates string `xml:"coordinates"`
}

// LineString represents a KML LineString element.
type LineString struct {
	Coordinates string `xml:"coordinates"`
}

// Polygon represents a KML Polygon element.
type Polygon struct {
	OuterBoundaryIs *Boundary `xml:"outerBoundaryIs>LinearRing,omitempty"`
	InnerBoundaryIs *Boundary `xml:"innerBoundaryIs>LinearRing,omitempty"`
}

// Boundary represents a LinearRing element in a KML Polygon.
type Boundary struct {
	Coordinates string `xml:"coordinates"`
}

// CoordinatesType represents a list of coordinates.
type CoordinatesType string

// TimeStamp represents a KML TimeStamp element.
type TimeStamp struct {
	When time.Time `xml:"when"`
}

// TimeSpan represents a KML TimeSpan element.
type TimeSpan struct {
	Begin time.Time `xml:"begin"`
	End   time.Time `xml:"end"`
}

// Metadata represents a KML Metadata element.
type Metadata struct {
	XMLName xml.Name `xml:"Metadata"`
	Any     string   `xml:",innerxml"`
}

// Schema represents a KML Schema element.
type Schema struct {
	Name         string        `xml:"name,attr"`
	SimpleFields []SimpleField `xml:"SimpleField"`
}

// SimpleField represents a KML SimpleField element.
type SimpleField struct {
	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
}

// Style represents a KML Style element.
type Style struct {
	XMLName      xml.Name      `xml:"Style"`
	LineStyle    *LineStyle    `xml:"LineStyle,omitempty"`
	PolyStyle    *PolyStyle    `xml:"PolyStyle,omitempty"`
	IconStyle    *IconStyle    `xml:"IconStyle,omitempty"`
	LabelStyle   *LabelStyle   `xml:"LabelStyle,omitempty"`
	BalloonStyle *BalloonStyle `xml:"BalloonStyle,omitempty"`
	ListStyle    *ListStyle    `xml:"ListStyle,omitempty"`
}

// LineStyle represents a KML LineStyle element.
type LineStyle struct {
	Color string  `xml:"color,omitempty"`
	Width float64 `xml:"width,omitempty"`
}

// PolyStyle represents a KML PolyStyle element.
type PolyStyle struct {
	Color   string `xml:"color,omitempty"`
	Fill    bool   `xml:"fill,omitempty"`
	Outline bool   `xml:"outline,omitempty"`
}

// IconStyle represents a KML IconStyle element.
type IconStyle struct {
	Color string  `xml:"color,omitempty"`
	Scale float64 `xml:"scale,omitempty"`
}

// LabelStyle represents a KML LabelStyle element.
type LabelStyle struct {
	Color string  `xml:"color,omitempty"`
	Scale float64 `xml:"scale,omitempty"`
}

// BalloonStyle represents a KML BalloonStyle element.
type BalloonStyle struct {
	BGColor   string `xml:"bgColor,omitempty"`
	TextColor string `xml:"textColor,omitempty"`
}

// ListStyle represents a KML ListStyle element.
type ListStyle struct {
	ListItemType string `xml:"listItemType,omitempty"`
}

// Region represents a KML Region element.
type Region struct {
	LatLonAltBox *LatLonAltBox `xml:"LatLonAltBox"`
	Lod          *Lod          `xml:"Lod"`
}

// LatLonAltBox represents a KML LatLonAltBox element.
type LatLonAltBox struct {
	North float64 `xml:"north"`
	South float64 `xml:"south"`
	East  float64 `xml:"east"`
	West  float64 `xml:"west"`
}

// Lod represents a KML Lod element.
type Lod struct {
	MinLodPixels  float64 `xml:"minLodPixels"`
	MaxLodPixels  float64 `xml:"maxLodPixels"`
	MinFadeExtent float64 `xml:"minFadeExtent"`
	MaxFadeExtent float64 `xml:"maxFadeExtent"`
}
