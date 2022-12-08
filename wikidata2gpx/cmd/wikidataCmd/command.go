package wikidataCmd

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"

	"github.com/d4l3k/go-pbzip2"

	"github.com/a1fred/wikidata2gpx.git/wikidata2gpx/gpxTools"
	"github.com/a1fred/wikidata2gpx.git/wikidata2gpx/utils"
	"github.com/a1fred/wikidata2gpx.git/wikidata2gpx/wikidataReader"
	"github.com/cheggaaa/pb/v3"
	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

var fout string
var lang string // https://www.wikidata.org/wiki/Help:Wikimedia_language_codes/lists/all
var cacheFilePath string
var gpxMetadataName string

func init() {
	WikidataCmd.Flags().StringVarP(&fout, "output", "o", "result.gpx", "result .gpx file")
	WikidataCmd.Flags().StringVarP(&cacheFilePath, "cache", "", ".wikidata2gpx_cache.bolt", "Cache file path, used for cache wikidata and nominatim api responses.")
	WikidataCmd.Flags().StringVarP(&lang, "lang", "l", "en", "Poi language")
	WikidataCmd.Flags().StringVarP(&gpxMetadataName, "gpx-metadata-name", "", "wikidata2gpx", "Gpx metadata name")
}

var WikidataCmd = &cobra.Command{
	Use:   "wikidata <wikidata-dump-file>",
	Short: "Generate gpx from wikidata dumps",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		main(args[0])
	},
}

func main(bz2DumpFilePath string) {
	start := time.Now()
	poiCounter := uint64(0)
	db, err := bolt.Open(cacheFilePath, 0666, nil)
	utils.ErrCheck(err)
	defer db.Close()

	entityNameExtractor := NewEntityNameExtractor(db, []byte(fmt.Sprintf("EntityNameExtractor-%s", lang)), lang)

	// Init writer
	poiChan := make(chan *gpxTools.Wpt, 5)
	var wg sync.WaitGroup

	exit := func(code int) {
		close(poiChan)
		wg.Wait()
		log.Printf("Done. %d saved. Took %s", poiCounter, time.Since(start))
		os.Exit(code)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			fmt.Printf("Interrupted: %v\n", sig.String())
			exit(1)
		}
	}()

	// Consumer
	fout, err := os.Create(fout)
	utils.ErrCheck(err)
	defer fout.Close()
	wg.Add(1)
	go gpxTools.WritePois(&wg, gpxMetadataName, poiChan, fout)

	// Producer
	fin, err := os.Open(bz2DumpFilePath)
	utils.ErrCheck(err)
	defer fin.Close()
	stat, err := fin.Stat()
	utils.ErrCheck(err)
	bar := pb.Full.Start64(stat.Size())
	barReader := bar.NewProxyReader(fin)
	// bufReader := barReader
	bufReader := bufio.NewReaderSize(barReader, 4096*5)
	var reader io.Reader

	if strings.HasSuffix(bz2DumpFilePath, ".bz2") {
		// https://github.com/d4l3k/go-pbzip2
		reader, err = pbzip2.NewReader(bufReader)
		utils.ErrCheck(err)
	} else if strings.HasSuffix(bz2DumpFilePath, ".gz") {
		reader, err = gzip.NewReader(bufReader)
		utils.ErrCheck(err)
	} else if strings.HasSuffix(bz2DumpFilePath, ".json") {
		reader = bufReader
	} else {
		log.Fatalln("Cant determine type of dump file")
	}

	wikidataReader.StreamEntitiesHandler(
		reader,
		func(e *wikidataReader.Entity) {
			for _, poi := range Entity2Poi(e, lang, entityNameExtractor) {
				poiCounter++
				poiChan <- poi
			}
		},
	)

	exit(0)
}
