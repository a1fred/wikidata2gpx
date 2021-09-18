package wikidataCmd

import (
	"github.com/a1fred/wikidata2gpx.git/wikidata2gpx/wikidataReader"
)

func extractLocalizedNameFromEntity(e *wikidataReader.Entity, lang string) string {
	nameElement, ok := e.Labels[lang]
	if !ok {
		return ""
	}
	return nameElement.Value
}

func extractLocalizedDescFromEntity(e *wikidataReader.Entity, lang string) string {
	descElement, ok := e.Descriptions[lang]
	if !ok {
		return ""
	}
	return descElement.Value
}
