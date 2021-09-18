package wikidataReader

import (
	"encoding/json"
	"fmt"
)

type Entity struct {
	Type         string                           `json:"type"`
	Datatype     string                           `json:"datatype"`
	Id           string                           `json:"id"`
	Labels       map[string]EntityLabelLanguage   `json:"labels"`
	Descriptions map[string]EntityLabelLanguage   `json:"descriptions"`
	Aliases      map[string][]EntityLabelLanguage `json:"aliases"`
	Claims       map[string][]EntityClaim         `json:"claims"`
	Sitelinks    map[string]EntitySitelink        `json:"sitelinks"`

	LastRevId uint64 `json:"lastrevid"`
}

func (e *Entity) Url() string {
	return fmt.Sprintf("https://www.wikidata.org/wiki/%s", e.Id)
}

type EntityLabelLanguage struct {
	Language string `json:"language"`
	Value    string `json:"value"`
}

type EntityClaim struct {
	Mainsnak        *EntityClaimMainsnak             `json:"mainsnak"`
	Type            string                           `json:"type"`
	Id              string                           `json:"id"`
	Rank            string                           `json:"rank"`
	Qualifiers      map[string][]EntityClaimMainsnak `json:"qualifiers"`
	QualifiersOrder []string                         `json:"qualifiers-order"`
	References      []EntityClaimReference           `json:"references"`
}

type EntityClaimMainsnak struct {
	Snaktype  string                       `json:"snaktype"`
	Property  string                       `json:"property"`
	Hash      string                       `json:"hash"`
	Datavalue EntityClaimMainsnakDatavalue `json:"datavalue"`
	Datatype  string                       `json:"datatype"`
}

type EntityClaimMainsnakDatavalue struct {
	Value json.RawMessage `json:"value"`
	Type  string          `json:"type"`
}

type EntityClaimReference struct {
	Hash       string                           `json:"hash"`
	Snaks      map[string][]EntityClaimMainsnak `json:"snaks"`
	SnaksOrder []string                         `json:"snaks-order"`
}

type EntitySitelink struct {
	Site   string   `json:"site"`
	Title  string   `json:"title"`
	Badges []string `json:"badges"`
}
