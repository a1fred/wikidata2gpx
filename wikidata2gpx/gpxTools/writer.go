package gpxTools

import (
	"fmt"
	"io"
	"sync"

	"github.com/a1fred/wikidata2gpx.git/wikidata2gpx/utils"
)

func mustWriteString(data string, to io.StringWriter) {
	_, err := to.WriteString(data)
	utils.ErrCheck(err)
}

func WritePois(wg *sync.WaitGroup, metadata_name string, c chan *Wpt, to io.StringWriter) {
	defer wg.Done()

	mustWriteString(fmt.Sprintf(`<?xml version='1.0' encoding='UTF-8' standalone='yes' ?>
<gpx version="1.1" creator="wikidata2gpx" xmlns="http://www.topografix.com/GPX/1/1" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.topografix.com/GPX/1/1 http://www.topografix.com/GPX/1/1/gpx.xsd">
  <metadata>
    <name>%s</name>
  </metadata>

`, escapeXml(metadata_name)), to)

	for poi := range c {
		mustWriteString(fmt.Sprintf("  %s\n", poi.ToXml()), to)
	}

	mustWriteString("</gpx>\n", to)
}
