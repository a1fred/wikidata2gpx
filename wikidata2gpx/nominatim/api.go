package nominatim

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

type NominatimApi struct {
	client *resty.Client
}

func NewNominatimApi() *NominatimApi {
	client := resty.New()
	// client.EnableTrace()

	client.
		SetTimeout(1 * time.Minute).
		SetRetryCount(15).
		SetRetryWaitTime(5 * time.Second)

	return &NominatimApi{
		client: client,
	}
}

type ReverseResponse struct {
	Error string `json:"error"`

	// OR
	PlaceId     int             `json:"place_id"`
	Licence     string          `json:"licence"`
	OsmType     string          `json:"osm_type"`
	OsmId       int             `json:"osm_id"`
	Lat         string          `json:"lat"`
	Lon         string          `json:"lon"`
	PlaceRank   int             `json:"place_rank"`
	Category    string          `json:"category"`
	Type        string          `json:"type"`
	Importance  float64         `json:"importance"`
	Addresstype string          `json:"addresstype"`
	DisplayName string          `json:"display_name"`
	Name        string          `json:"name"`
	Address     json.RawMessage `json:"address"`
	Boundingbox []string        `json:"boundingbox"`
}

func (r *ReverseResponse) GetReverseResponseAddress() (*ReverseResponseAddress, error) {
	addr := ReverseResponseAddress{}
	err := json.Unmarshal(r.Address, &addr)
	return &addr, err
}

type ReverseResponseAddress struct {
	State       string `json:"state"`
	CountryCode string `json:"country_code"`
}

func (n *NominatimApi) Reverse(lat, lon float64) (*ReverseResponse, error) {
	/*
		https://nominatim.openstreetmap.org/reverse?format=jsonv2&lat=-74.695000&lon=164.114000
	*/
	resp := ReverseResponse{}
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=jsonv2&lat=%f&lon=%f", lat, lon)

	r, err := n.client.R().
		SetResult(&resp).
		Get(url)

	if err != nil {
		return nil, fmt.Errorf("%s, %s: %s", url, r.String(), err)
	}

	if r.StatusCode() != 200 {
		return nil, fmt.Errorf("error fetching %f,%f: %s", lat, lon, r.String())
	}

	return &resp, nil
}
