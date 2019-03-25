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

	"golang.org/x/xerrors"

	"go.felesatra.moe/dlsite"
	"go.felesatra.moe/dlsite/cache"
	"go.felesatra.moe/dlsite/hvdb"
)

type nullCache struct{}

func (nullCache) Close() error {
	return nil
}

func (nullCache) Get(r dlsite.RJCode) (*dlsite.Work, error) {
	return nil, xerrors.Errorf("get %s from NullCache", r)
}

func (nullCache) Put(w *dlsite.Work) error {
	return nil
}

// Cache defines the interface for cache objects to pass to Fetch.
// The semantics of the methods should match cache.Cache.
type Cache interface {
	Close() error
	Get(dlsite.RJCode) (*dlsite.Work, error)
	Put(*dlsite.Work) error
}

// DefaultCache returns a sane default Cache object.  This cache is
// backed by a database stored at a path respecting XDG directory
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
		log.Printf("Error fetching from DLSite: %s", err)
		hw, err := hvdb.Fetch(r)
		if err != nil {
			return nil, xerrors.Errorf("fetch from HVDB: %w", err)
		}
		w = convertWork(hw)
	}
	if err := c.Put(w); err != nil {
		log.Printf("Failed to cache work: %s", err)
	}
	return w, nil
}

func convertWork(w *hvdb.Work) *dlsite.Work {
	return &dlsite.Work{
		RJCode: w.RJCode,
		Name:   w.Title,
		Maker:  w.Circle,
	}
}
