package kml

type GeoPoint struct {
	Latitude  float64
	Longitude float64
	Altitude  float64
}

type PolygonData struct {
	Name        string
	Coordinates []GeoPoint
}
