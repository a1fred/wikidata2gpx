package wikidataReader

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type WikidataApi struct {
	client *resty.Client
}

func NewWikidataApi() *WikidataApi {
	client := resty.New()
	// client.EnableTrace()

	return &WikidataApi{
		client: client,
	}
}

type EntityDataResponse struct {
	Entities map[string]*Entity `json:"entities"`
}

func (w *WikidataApi) EntityData(id string) (*EntityDataResponse, error) {
	resp := EntityDataResponse{}

	r, err := w.client.R().
		SetResult(&resp).
		Get(fmt.Sprintf("https://www.wikidata.org/wiki/Special:EntityData/%s.json", id))

	if err != nil {
		return nil, err
	}

	if r.StatusCode() != 200 {
		return nil, fmt.Errorf("error fetching %s: %s", id, r.String())
	}

	return &resp, nil
}
