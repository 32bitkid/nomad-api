package main

import (
	"encoding/xml"
	"io"
	"math"
)

type Gpx interface {
	TrackPoints() []TrackPoint
}

type gpx struct {
	Creator string           `xml:"creator,attr"`
	Time    string           `xml:"metadata>time"`
	Title   string           `xml:"trk>name"`
	Trkpts  []*xmlTrackPoint `xml:"trk>trkseg>trkpt"`
}

func (g *gpx) TrackPoints() []TrackPoint {
	trackPoints := make([]TrackPoint, 0, len(g.Trkpts))
	for _, p := range g.Trkpts {
		trackPoints = append(trackPoints, p)
	}
	return trackPoints
}

type xmlTrackPoint struct {
	Lat float64 `xml:"lat,attr"`
	Lon float64 `xml:"lon,attr"`
	Ele float32 `xml:"ele"`
}

type TrackPoint interface {
	Latitude() float64
	Longitude() float64
	Elevation() float32
	DistanceTo(TrackPoint) float64
}

const EarthRadius float64 = 6371

var toRad = func(val float64) float64 { return val * math.Pi / 180 }

func FromXml(file io.Reader) (Gpx, error) {
	decoder := xml.NewDecoder(file)

	g := &gpx{}
	e := decoder.Decode(g)
	if e != nil {
		return nil, e
	}

	return g, nil
}

func (p *xmlTrackPoint) Latitude() float64 {
	return p.Lat
}

func (p *xmlTrackPoint) Longitude() float64 {
	return p.Lon
}

func (p *xmlTrackPoint) Elevation() float32 {
	return p.Ele
}

func (start *xmlTrackPoint) DistanceTo(end TrackPoint) float64 {
	return DistanceBetween(start, end)
}

func DistanceBetween(start, end TrackPoint) float64 {
	dLat := end.Latitude() - start.Latitude()
	dLon := end.Longitude() - start.Longitude()

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(start.Latitude())*math.Cos(end.Latitude())*
			math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return EarthRadius * c
}
