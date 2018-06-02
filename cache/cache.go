/*
Package cache provides caching for DLSite work information.
*/
package cache

import (
	"bytes"
	"encoding/gob"
	"path/filepath"

	"github.com/coreos/bbolt"
	"github.com/pkg/errors"
	"go.felesatra.moe/dlsite"
	"go.felesatra.moe/xdg"
)

// Cache is a cache for DLSite work information.
type Cache struct {
	db *bolt.DB
}

// OpenDefault opens the default cache using XDG directory
// configuration.
func OpenDefault() (*Cache, error) {
	p := filepath.Join(xdg.CacheHome(), "go.felesatra.moe_dlsite.db")
	return Open(p)
}

// Open opens and returns a Cache object.
func Open(p string) (*Cache, error) {
	db, err := bolt.Open(p, 0600, nil)
	if err != nil {
		return nil, errors.Wrap(err, "open DLSite cache")
	}
	return &Cache{db: db}, nil
}

// Close closes the Cache.
func (c *Cache) Close() error {
	return c.db.Close()
}

var bucket []byte = []byte("dlsite")

// Get returns the work with the RJCode in the Cache.
func (c *Cache) Get(r dlsite.RJCode) (*dlsite.Work, error) {
	var w *dlsite.Work
	err := c.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		if b == nil {
			return errors.Errorf("bucket missing")
		}
		d := b.Get(encodeRJCode(r))
		if d == nil {
			return errors.Errorf("get %s: missing", c)
		}
		var err error
		w, err = decodeWork(d)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return w, nil
}

// Put inserts the work into the Cache.
func (c *Cache) Put(w *dlsite.Work) error {
	return c.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		if b == nil {
			var err error
			b, err = tx.CreateBucket(bucket)
			if err != nil {
				return err
			}
		}
		err := b.Put(encodeRJCode(w.RJCode), encodeWork(w))
		return err
	})
}

func encodeRJCode(c dlsite.RJCode) []byte {
	return []byte(c)
}

func encodeWork(w *dlsite.Work) []byte {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	enc.Encode(w)
	return b.Bytes()
}

func decodeWork(b []byte) (*dlsite.Work, error) {
	dec := gob.NewDecoder(bytes.NewReader(b))
	var w dlsite.Work
	err := dec.Decode(&w)
	if err != nil {
		return nil, errors.Wrap(err, "decode dlsite.Work")
	}
	return &w, nil
}
