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
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"

	"go.felesatra.moe/dlsite/v2/codes"
	"go.felesatra.moe/dlsite/v2/internal/caching"
	"go.felesatra.moe/dlsite/v2/internal/dlsite"
	"go.felesatra.moe/dlsite/v2/internal/hvdb"
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
	cachePath string
	cmap      *caching.Map
}

// NewFetcher creates a new Fetcher.
// The Fetcher scrapes DLSite and HVDB for information.
// The Fetcher keeps a cache in the user's XDG cache directory.
func NewFetcher(o ...FetcherOption) (*Fetcher, error) {
	f := &Fetcher{
		cachePath: filepath.Join(userCacheDir(), "go.felesatra.moe_dlsite_v2.cache"),
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

// Close closes resources used by the fetcher.
func (f *Fetcher) Close() error {
	if err := f.cmap.Close(); err != nil {
		return fmt.Errorf("dlsite: %s", err)
	}
	return nil
}

// FetchWork fetches information for a DLSite work.
//
// The FetchWorkOption argument is deprecated; use the
// FetchWorkDirectly method instead.
func (f *Fetcher) FetchWork(c codes.WorkCode, o ...FetchWorkOption) (*Work, error) {
	opts := mergeOptions(o...)
	if opts.ignoreCache {
		return f.FetchWorkDirectly(c)
	}
	if w := f.getCached(c); w != nil {
		return w, nil
	}
	return f.FetchWorkDirectly(c)
}

// FetchWorkDirectly fetches information for a DLSite work, bypassing
// any cache.
// The newly fetched information will still be cached for future
// FetchWork calls.
func (f *Fetcher) FetchWorkDirectly(c codes.WorkCode) (*Work, error) {
	w, err := f.fetchWork(c)
	if err != nil {
		return nil, fmt.Errorf("dlsite: %s", err)
	}
	f.putCached(c, w)
	return w, nil
}

func (*Fetcher) fetchWork(c codes.WorkCode) (*Work, error) {
	w := &Work{}
	var ok bool
	dw, err := dlsite.FetchWork(codes.RJCode(c))
	if err != nil {
		log.Printf("dlsite: %s", err)
	} else {
		ok = true
		fillWorkFromDLSite(w, dw)
	}
	if dw == nil || len(dw.WorkFormats) == 0 || dw.WorkFormats[0] == "ボイス・ASMR" {
		if err := fillWorkFromHVDB(w, codes.RJCode(c)); err != nil {
			log.Printf("dlsite: %s", err)
		} else {
			ok = true
		}
	}
	if !ok {
		return nil, fmt.Errorf("fetch work %s: all methods failed", c)
	}
	w.CVs = removeDupes(w.CVs)
	return w, nil
}

func (f *Fetcher) getCached(c codes.WorkCode) *Work {
	if f.cmap == nil {
		return nil
	}
	var w *Work
	f.cmap.Get(c, &w)
	return w
}

func (f *Fetcher) putCached(c codes.WorkCode, w *Work) {
	if f.cmap == nil {
		return
	}
	f.cmap.Put(c, w)
}

func fillWorkFromDLSite(w *Work, dw *dlsite.Work) {
	if dw.Code != "" {
		w.Code = codes.WorkCode(dw.Code)
	}
	if dw.Title != "" {
		w.Title = dw.Title
	}
	if dw.Circle != "" {
		w.Circle = dw.Circle
	}
	if dw.Series != "" {
		w.Series = dw.Series
	}
	if len(dw.Seiyuu) > 0 {
		w.CVs = dw.Seiyuu
	}
	if dw.Description != "" {
		w.Description = dw.Description
	}
}

func fillWorkFromHVDB(w *Work, c codes.RJCode) error {
	hw, err := hvdb.FetchWork(c)
	if err != nil {
		return err
	}
	if hw.Code != "" {
		w.Code = codes.WorkCode(hw.Code)
	}
	if w.Title == "" && hw.Title != "" {
		w.Title = hw.Title
	}
	if hw.EnglishTitle != "" {
		w.EnglishTitle = hw.EnglishTitle
	}
	if w.Circle == "" && hw.Circle != "" {
		w.Circle = hw.Circle
	}
	for _, v := range hw.CVs {
		if v == "N/A" {
			continue
		}
		w.CVs = append(w.CVs, v)
	}
	w.Tags = append(w.Tags, hw.Tags...)
	w.SFW = hw.SFW
	return nil
}

func removeDupes(v []string) []string {
	if len(v) < 2 {
		return v
	}
	sort.Strings(v)
	i := 1
	for j := 1; j < len(v); j++ {
		if v[i-1] != v[j] {
			v[i] = v[j]
			i++
		}
	}
	return v[:i]
}

func userCacheDir() string {
	d := os.Getenv("XDG_CACHE_HOME")
	if d == "" {
		d = filepath.Join(os.Getenv("HOME"), ".cache")
	}
	return d
}
