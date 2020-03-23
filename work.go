// Copyright (C) 2020 Allen Li
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

// Package dlsite provides access to DLSite work information.
package dlsite

import (
	"encoding/gob"
	"errors"
	"fmt"
	"path/filepath"

	"go.felesatra.moe/xdg"

	"go.felesatra.moe/dlsite/v2/codes"
	"go.felesatra.moe/dlsite/v2/internal/caching"
)

func init() {
	gob.Register((*Work)(nil))
}

// A Work contains information about a DLSite work.
type Work struct {
	Code         codes.WorkCode
	Title        string
	EnglishTitle string
	Circle       string
	Series       string
	Description  string
	CVs          []string
	Tags         []string
	SFW          bool
}

// A Fetcher can fetch DLSite work information.
type Fetcher struct {
	cachePath  string
	cmap       *caching.Map
	forceFresh bool
}

// NewFetcher creates a new Fetcher.
// The Fetcher scrapes DLSite and HVDB for information.
// The Fetcher keeps a cache in the user's XDG cache directory.
func NewFetcher(o ...FetcherOption) (*Fetcher, error) {
	f := &Fetcher{
		cachePath: filepath.Join(xdg.CacheHome(), "go.felesatra.moe_dlsite_v2.cache"),
	}
	for _, o := range o {
		o.apply(f)
	}
	if f.cachePath != "" {
		cmap, err := caching.Open(f.cachePath)
		if err != nil {
			return nil, fmt.Errorf("dlsite: %s", err)
		}
		f.cmap = cmap
	}
	return f, nil
}

// FetchWork fetches information for a DLSite work.
func (f *Fetcher) FetchWork(c codes.WorkCode) (*Work, error) {
	if w := f.getCached(c); w != nil {
		return w, nil
	}
	w := &Work{}
	// TODO: Validate RJ code
	if err := fillWorkFromHVDB(w, codes.RJCode(c)); err != nil {
		return nil, fmt.Errorf("dlsite: %s", err)
	}
	if err := fillWorkFromDLSite(w, codes.RJCode(c)); err != nil {
		return nil, fmt.Errorf("dlsite: %s", err)
	}
	return w, nil
}

// FlushCache flushes the cache, reading and writing changes from the cache file.
// An error is returned if the cache path is empty.
func (f *Fetcher) FlushCache() error {
	if f.cmap == nil {
		return errors.New("dlsite flush cache: no cache")
	}
	return f.cmap.Flush()
}

func (f *Fetcher) getCached(c codes.WorkCode) *Work {
	if f.cmap == nil || f.forceFresh {
		return nil
	}
	w, ok := f.cmap.Get(c).(*Work)
	if !ok {
		return nil
	}
	return w
}

func (f *Fetcher) putCached(c codes.WorkCode, w *Work) {
	if f.cmap == nil {
		return
	}
	f.cmap.Put(c, w)
}

// A FetcherOption can be passed to NewFetcher to configure Fetcher creation.
type FetcherOption interface {
	apply(*Fetcher)
	fetcherOption()
}

type cacheOption struct {
	path string
}

func (o cacheOption) apply(f *Fetcher) {
	f.cachePath = o.path
}

func (cacheOption) fetcherOption() {}

// CachePath sets the cache path of the Fetcher.
// If path is empty, no cache file is used.
func CachePath(path string) FetcherOption {
	return cacheOption{path}
}

type freshOption struct{}

func (o freshOption) apply(f *Fetcher) {
	f.forceFresh = true
}

func (freshOption) fetcherOption() {}

// ForceFresh makes the Fetcher ignore the cache when getting work
// information.
// Updated work information is still added to the cache if a cache
// path is provided.
func ForceFresh() FetcherOption {
	return freshOption{}
}
