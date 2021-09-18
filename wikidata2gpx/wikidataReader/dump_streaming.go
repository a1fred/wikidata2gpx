package wikidataReader

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/a1fred/wikidata2gpx.git/wikidata2gpx/utils"
)

func StreamEntitiesChan(dumpFileReader io.Reader, out chan *Entity) {
	jsonDecoder := json.NewDecoder(dumpFileReader)
	jsonDecoder.DisallowUnknownFields()

	// read open bracket
	_, err := jsonDecoder.Token()
	utils.ErrCheck(err)

	// while the array contains values
	for jsonDecoder.More() {
		var e Entity
		err := jsonDecoder.Decode(&e)
		if err != nil {
			log.Fatal(fmt.Sprintf("%s at offset %d", err, jsonDecoder.InputOffset()))
		}

		out <- &e
	}

	// read closing bracket
	_, err = jsonDecoder.Token()
	utils.ErrCheck(err)

}

func StreamEntitiesHandler(dumpFileReader io.Reader, handler func(*Entity)) {
	c := make(chan *Entity)

	go func() {
		StreamEntitiesChan(dumpFileReader, c)
		close(c)
	}()

	for e := range c {
		handler(e)
	}
}
