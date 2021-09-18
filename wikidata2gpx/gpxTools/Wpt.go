package gpxTools

import (
	"encoding/xml"
	"strconv"

	"github.com/a1fred/wikidata2gpx.git/wikidata2gpx/utils"
)

// http://www.topografix.com/GPX/1/1/#type_wptType
type Wpt struct {
	XMLName xml.Name `xml:"wpt"`

	Lat string `xml:"lat,attr"` // xsd:decimal
	Lon string `xml:"lon,attr"` // xsd:decimal

	Name string `xml:"name"`           // xsd:string
	Desc string `xml:"desc,omitempty"` // xsd:string

	Link []string `xml:"link,omitempty"` // xsd:string

	Cmt string `xml:"cmt,omitempty"` // xsd:string
	Src string `xml:"src,omitempty"` // xsd:string

	Type string `xml:"type,omitempty"` // xsd:string
}

func (p *Wpt) ToXml() string {
	encoded, err := xml.Marshal(p)
	if err != nil {
		panic(err)
	}
	return string(encoded)
}

func (p *Wpt) MustLatFloat64() float64 {
	val, err := strconv.ParseFloat(p.Lat, 64)
	utils.ErrCheck(err)
	return val
}

func (p *Wpt) MustLonFloat64() float64 {
	val, err := strconv.ParseFloat(p.Lon, 64)
	utils.ErrCheck(err)
	return val
}
