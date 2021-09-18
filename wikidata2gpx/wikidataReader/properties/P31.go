package properties

import (
	"encoding/json"
	"fmt"
)

type P31 struct {
	Id         string `json:"id"`
	NumericId  int    `json:"numeric-id"`
	EntityType string `json:"entity-type"`
}

func NewP31(claimDataValue json.RawMessage) *P31 {
	if string(claimDataValue) == "" {
		return nil
	}

	prop := P31{}
	err := json.Unmarshal(claimDataValue, &prop)
	if err != nil {
		panic(fmt.Sprintf("json '%s' parse error: %s", string(claimDataValue), err))
	}
	return &prop
}
