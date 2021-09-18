package properties

import (
	"encoding/json"
	"fmt"
)

type P625 struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  float64 `json:"altitude"`
	Precision float64 `json:"precision"`
	Globe     string  `json:"globe"`
}

func NewP625(claimDataValue json.RawMessage) *P625 {
	if string(claimDataValue) == "" {
		return nil
	}

	prop := P625{}
	err := json.Unmarshal(claimDataValue, &prop)
	if err != nil {
		panic(fmt.Sprintf("json '%s' parse error: %s", string(claimDataValue), err))
	}
	return &prop
}
