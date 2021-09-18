package filterCmd

import (
	"encoding/json"
	"fmt"

	"github.com/a1fred/wikidata2gpx.git/wikidata2gpx/nominatim"
	"github.com/a1fred/wikidata2gpx.git/wikidata2gpx/utils"
	bolt "go.etcd.io/bbolt"
)

type ReverseGeocoder struct {
	nominatimApi *nominatim.NominatimApi
	db           *bolt.DB // Caching database
	bucketName   []byte
}

func NewReverseGeocoder(db *bolt.DB, dbBucketName []byte, nominatimApi *nominatim.NominatimApi) *ReverseGeocoder {
	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(dbBucketName)
		return err
	})
	utils.ErrCheck(err)

	return &ReverseGeocoder{
		nominatimApi: nominatimApi,
		db:           db,
		bucketName:   dbBucketName,
	}
}

type CacheItem struct {
	CountryCode string
	State       string
}

func (n *ReverseGeocoder) Reverse(lat, lon float64) *CacheItem {
	cacheKey := []byte(fmt.Sprintf("%f-%f", lat, lon))
	cacheItem := CacheItem{}

	var cached []byte
	err := n.db.View(func(tx *bolt.Tx) error {
		cached = tx.Bucket(n.bucketName).Get(cacheKey)
		return nil
	})
	utils.ErrCheck(err)

	if cached != nil {
		err := json.Unmarshal(cached, &cacheItem)
		utils.ErrCheck(err)
		return &cacheItem
	}

	response, err := n.nominatimApi.Reverse(lat, lon)
	utils.ErrCheck(err)

	if response.Error != "" {
		fmt.Printf("nominatim error: coords %f, %f: %s\n", lat, lon, response.Error)
		response.Address = json.RawMessage("{}")
	}

	addr, err := response.GetReverseResponseAddress()
	if err != nil {
		fmt.Printf("nominatim wrong address value: %s\n", err)
		addr = &nominatim.ReverseResponseAddress{}
	}

	cacheItem = CacheItem{
		CountryCode: addr.CountryCode,
		State:       addr.State,
	}

	cached, err = json.Marshal(cacheItem)
	utils.ErrCheck(err)

	err = n.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(n.bucketName)
		err := b.Put(cacheKey, cached)
		return err
	})
	utils.ErrCheck(err)

	return &cacheItem
}
