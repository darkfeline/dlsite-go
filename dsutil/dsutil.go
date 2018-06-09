// Copyright (C) 2018 Allen Li
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
Package dsutil implements a simplified API for looking up DLSite work
information.
*/
package dsutil

import (
	"log"

	"github.com/pkg/errors"
	"go.felesatra.moe/dlsite"
	"go.felesatra.moe/dlsite/cache"
)

type nullCache struct{}

func (_ nullCache) Close() error {
	return nil
}

func (_ nullCache) Get(r dlsite.RJCode) (*dlsite.Work, error) {
	return nil, errors.Errorf("get %s from NullCache", r)
}

func (_ nullCache) Put(w *dlsite.Work) error {
	return nil
}

// Cache defines the interface for cache objects to pass to Fetch.
// The semantics of the methods should match dlsite.Cache.
type Cache interface {
	Close() error
	Get(dlsite.RJCode) (*dlsite.Work, error)
	Put(*dlsite.Work) error
}

// DefaultCache returns a sane default Cache object.  This cache is
// backed by a database stored at a path respecting XDG database
// specifications, but falls back to a null object if opening the file
// fails.  Make sure to defer a call to Close.
func DefaultCache() Cache {
	c, err := cache.OpenDefault()
	if err != nil {
		log.Printf("Error opening cache: %s", err)
		return nullCache{}
	}
	return c
}

// Fetch fetches work information using a persistent cache.
func Fetch(c Cache, r dlsite.RJCode) (*dlsite.Work, error) {
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
