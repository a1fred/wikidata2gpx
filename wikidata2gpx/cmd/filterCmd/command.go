package filterCmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/a1fred/wikidata2gpx.git/wikidata2gpx/gpxTools"
	"github.com/a1fred/wikidata2gpx.git/wikidata2gpx/nominatim"
	"github.com/a1fred/wikidata2gpx.git/wikidata2gpx/utils"
	"github.com/cheggaaa/pb/v3"
	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

var fout string
var cacheFilePath string
var countryCode string // https://en.wikipedia.org/wiki/ISO_3166-1
var gpxMetadataName string

func init() {
	dirname, err := os.UserHomeDir()
	utils.ErrCheck(err)

	FilterCmd.Flags().StringVarP(&fout, "output", "o", "result.gpx", "result .gpx file")
	FilterCmd.Flags().StringVarP(&cacheFilePath, "cache", "", path.Join(dirname, ".wikidata2gpx_cache.bolt"), "Cache file path, used for cache wikidata and nominatim api responses.")
	FilterCmd.Flags().StringVarP(&countryCode, "country-code", "", "", "Generate only for country code (ISO_3166-1), all countries if empty (default)")
	FilterCmd.Flags().StringVarP(&gpxMetadataName, "gpx-metadata-name", "", "wikidata2gpx", "Gpx metadata name")
}

var FilterCmd = &cobra.Command{
	Use:   "filter <gpx-file.gpx>",
	Short: "Filter gpx file points",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		main(args[0])
	},
}

func main(gpxInFilePath string) {
	start := time.Now()
	poiCounter := uint64(0)
	db, err := bolt.Open(cacheFilePath, 0666, nil)
	utils.ErrCheck(err)
	defer db.Close()

	reverseGeocoder := NewReverseGeocoder(db, []byte("CachedNominatimApi"), nominatim.NewNominatimApi())

	// Init writer
	poiChan := make(chan *gpxTools.Wpt, 3)
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
	fin, err := os.Open(gpxInFilePath)
	utils.ErrCheck(err)
	defer fin.Close()
	stat, err := fin.Stat()
	utils.ErrCheck(err)
	bar := pb.Full.Start64(stat.Size())
	barReader := bar.NewProxyReader(fin)
	bufReader := bufio.NewReader(barReader)

	gpxTools.StreamWpt(
		bufReader,
		func(w *gpxTools.Wpt) {
			// Filter countryCode
			if countryCode != "" {
				nominatimResponse := reverseGeocoder.Reverse(w.MustLatFloat64(), w.MustLonFloat64())
				if !strings.EqualFold(nominatimResponse.CountryCode, countryCode) {
					return
				}
			}

			poiCounter++
			poiChan <- w
		},
	)

	exit(0)
}
