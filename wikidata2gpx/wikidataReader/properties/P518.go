package properties

import (
	"encoding/json"
	"fmt"
)

type P518 struct {
	Id         string `json:"id"`
	NumericId  int    `json:"numeric-id"`
	EntityType string `json:"entity-type"`
}

func NewP518(claimDataValue json.RawMessage) *P518 {
	if string(claimDataValue) == "" {
		return nil
	}

	prop := P518{}
	err := json.Unmarshal(claimDataValue, &prop)
	if err != nil {
		panic(fmt.Sprintf("json '%s' parse error: %s", string(claimDataValue), err))
	}
	return &prop
}
