package wikidataCmd

import (
	"fmt"
	"strings"

	"github.com/a1fred/wikidata2gpx.git/wikidata2gpx/gpxTools"
	"github.com/a1fred/wikidata2gpx.git/wikidata2gpx/wikidataReader"
	"github.com/a1fred/wikidata2gpx.git/wikidata2gpx/wikidataReader/properties"
)

func Entity2Poi(
	e *wikidataReader.Entity,
	lang string,
	entityNameExtractor *EntityNameExtractor,
) []*gpxTools.Wpt {
	results := make([]*gpxTools.Wpt, 0)

	for _, p625claim := range wikidataReader.ExtractClaimsFromEntity(e, "P625") {
		p625 := properties.NewP625(p625claim.Mainsnak.Datavalue.Value)
		// Decode result may be nil
		if p625 == nil {
			continue
		}

		// Instance of
		p31 := &properties.P31{}
		for _, p31claim := range wikidataReader.ExtractClaimsFromEntity(e, "P31") {
			p31 = properties.NewP31(p31claim.Mainsnak.Datavalue.Value)
			break
		}

		entityName := extractLocalizedNameFromEntity(e, lang)
		entityDesc := extractLocalizedDescFromEntity(e, lang)
		entityWikiUrl := wikidataReader.GetSitelinkWikiUrl(e, lang)

		// Ignore empty name or empty url points
		if entityName == "" || entityWikiUrl == "" {
			continue
		}

		// Ignore non-earth points
		if p625.Globe != "" && p625.Globe != "http://www.wikidata.org/entity/Q2" {
			continue
		}

		// Type id to local name
		entityType := ""
		if p31.Id != "" {
			entityType = entityNameExtractor.ExtractName(p31.Id)
		}

		// Append coord qualifiers to point name
		p625Qualifiers := make([]string, 0)
		for _, qualifierSlice := range p625claim.Qualifiers {
			for _, qualifier := range qualifierSlice {
				if qualifier.Property == "P518" {
					p518 := properties.NewP518(qualifier.Datavalue.Value)
					p518LocalName := entityNameExtractor.ExtractName(p518.Id)
					if p518LocalName != "" {
						p625Qualifiers = append(p625Qualifiers, p518LocalName)
					}
				}
			}
		}
		if len(p625Qualifiers) != 0 {
			entityName = fmt.Sprintf(
				"%s (%s)",
				entityName,
				strings.Join(p625Qualifiers, ", "),
			)
		}

		results = append(
			results,
			&gpxTools.Wpt{
				Lat: fmt.Sprintf("%.6f", p625.Latitude),
				Lon: fmt.Sprintf("%.6f", p625.Longitude),

				Name: entityName,
				Desc: strings.TrimSpace(fmt.Sprintf("%s\n%s", entityDesc, entityWikiUrl)), // Join url because fucking maps software ignoring link tag ;(
				Link: []string{entityWikiUrl},

				Src:  e.Url(),
				Type: entityType,
			},
		)
	}
	return results
}
