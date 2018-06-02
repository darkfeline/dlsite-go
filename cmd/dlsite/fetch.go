package main

import (
	"log"

	"github.com/pkg/errors"
	"go.felesatra.moe/dlsite"
	"go.felesatra.moe/dlsite/cache"
)

type NullCache struct{}

func (_ NullCache) Close() error {
	return nil
}

func (_ NullCache) Get(r dlsite.RJCode) (*dlsite.Work, error) {
	return nil, errors.Errorf("get %s from NullCache", r)
}

func (_ NullCache) Put(w *dlsite.Work) error {
	return nil
}

type Cache interface {
	Close() error
	Get(dlsite.RJCode) (*dlsite.Work, error)
	Put(*dlsite.Work) error
}

func defaultCache() Cache {
	c, err := cache.OpenDefault()
	if err != nil {
		log.Printf("Error opening cache: %s", err)
		return NullCache{}
	}
	return c
}

func fetch(c Cache, r dlsite.RJCode) (*dlsite.Work, error) {
	w, err := c.Get(r)
	if err == nil {
		return w, nil
	}
	w, err = dlsite.Fetch(r)
	if err != nil {
		return nil, errors.Wrap(err, "fetch from DLSite")
	}
	if err := c.Put(w); err != nil {
		log.Printf("Failed to cache work: %s", err)
	}
	return w, nil
}
