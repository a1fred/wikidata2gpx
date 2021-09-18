package gpxTools

import (
	"encoding/xml"
	"io"

	"github.com/a1fred/wikidata2gpx.git/wikidata2gpx/utils"
)

func StreamWpt(reader io.Reader, handler func(w *Wpt)) {
	xmlDecoder := xml.NewDecoder(reader)

	_, err := xmlDecoder.Token()
	utils.ErrCheck(err)

	for {
		tok, err := xmlDecoder.Token()
		if tok == nil || err == io.EOF {
			// EOF means we're done.
			break
		} else if err != nil {
			utils.ErrCheck(err)
		}

		switch ty := tok.(type) {
		case xml.StartElement:
			if ty.Name.Local == "wpt" {
				w := Wpt{}
				err = xmlDecoder.DecodeElement(&w, &ty)
				utils.ErrCheck(err)
				handler(&w)
			}
		default:
		}
	}
}
