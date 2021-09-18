package wikidataCmd

import (
	"encoding/json"
	"log"

	"github.com/a1fred/wikidata2gpx.git/wikidata2gpx/utils"
	"github.com/a1fred/wikidata2gpx.git/wikidata2gpx/wikidataReader"
	bolt "go.etcd.io/bbolt"
)

type entityLanguages struct {
	Labels map[string]wikidataReader.EntityLabelLanguage `json:"labels"`
}

type EntityNameExtractor struct {
	api        *wikidataReader.WikidataApi
	lang       string
	db         *bolt.DB // Caching database
	bucketName []byte
}

func NewEntityNameExtractor(db *bolt.DB, dbBucketName []byte, lang string) *EntityNameExtractor {
	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(dbBucketName)
		return err
	})
	utils.ErrCheck(err)

	return &EntityNameExtractor{
		api:        wikidataReader.NewWikidataApi(),
		lang:       lang,
		db:         db,
		bucketName: dbBucketName,
	}
}

// Caches all languages to future usage
// Can be empty string
func (c *EntityNameExtractor) ExtractName(entityId string) string {
	var result = entityLanguages{}
	var cached []byte

	c.db.View(func(tx *bolt.Tx) error {
		cached = tx.Bucket(c.bucketName).Get([]byte(entityId))
		return nil
	})

	if cached != nil {
		err := json.Unmarshal(cached, &result)
		utils.ErrCheck(err)
	} else {
		resp, err := c.api.EntityData(entityId)
		utils.ErrCheck(err)

		if len(resp.Entities) != 1 {
			log.Fatalf("cant find entity: %s, response contains %d elements", entityId, len(resp.Entities))
		}

		var entity *wikidataReader.Entity
		for k := range resp.Entities {
			entity = resp.Entities[k]
			break
		}

		result.Labels = entity.Labels

		err = c.db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket(c.bucketName)
			data, err := json.Marshal(&result)
			utils.ErrCheck(err)
			err = b.Put([]byte(entityId), data)
			return err
		})
		utils.ErrCheck(err)
	}

	nameElement, ok := result.Labels[lang]
	if !ok {
		return ""
	}
	return nameElement.Value
}
