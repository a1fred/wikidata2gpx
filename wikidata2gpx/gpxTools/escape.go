package gpxTools

import "strings"

func escapeXml(data string) string {
	data = strings.ReplaceAll(data, "\"", `\\\"`)       // escapeQuotes
	data = strings.ReplaceAll(data, "\n", "&#13;&#10;") // escapeNewlines
	return data
}
