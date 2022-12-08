package wikidataReader

import (
	"fmt"
	"net/url"
)

func getSitelinkTitle(e *Entity, findName string) string {
	for sitelinkName, sitelinkValue := range e.Sitelinks {
		if sitelinkName == findName {
			return url.PathEscape(sitelinkValue.Title)
		}
	}

	return ""

}

func GetSitelinkWikiUrl(e *Entity, lang string) string {
	// TODO: add all or rewrite
	switch lang {
	case "ru":
		title := getSitelinkTitle(e, "ruwiki")
		if title == "" {
			return ""
		}

		return fmt.Sprintf("https://ru.wikipedia.org/wiki/%s", title)
	case "en":
		title := getSitelinkTitle(e, "enwiki")
		if title == "" {
			return ""
		}

		return fmt.Sprintf("https://en.wikipedia.org/wiki/%s", title)
	}

	return ""
}
