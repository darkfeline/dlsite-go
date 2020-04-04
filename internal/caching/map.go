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

// Package caching implements a key value caching store.
//
// The stored values must be registered with gob.Register.
package caching

import (
	"compress/gzip"
	"encoding/gob"
	"fmt"
	"io"
	"os"

	"go.felesatra.moe/dlsite/v2/codes"
	"go.felesatra.moe/dlsite/v2/internal/lockedfile"
)

type Map struct {
	path     string
	cached   map[codes.WorkCode]interface{}
	modified map[codes.WorkCode]interface{}
}

func Open(p string) (*Map, error) {
	m := &Map{
		path:     p,
		cached:   make(map[codes.WorkCode]interface{}),
		modified: make(map[codes.WorkCode]interface{}),
	}
	f, err := lockedfile.Open(p)
	if err != nil {
		if os.IsNotExist(err) {
			return m, nil
		}
		return nil, fmt.Errorf("open caching map: %s", err)
	}
	defer f.Close()
	if err := m.readFrom(f); err != nil {
		return nil, fmt.Errorf("open caching map: %s", err)
	}
	return m, nil
}

func (m *Map) Get(c codes.WorkCode) interface{} {
	w, ok := m.modified[c]
	if ok {
		return w
	}
	w, ok = m.cached[c]
	if ok {
		return w
	}
	return nil
}

func (m *Map) Put(c codes.WorkCode, w interface{}) {
	m.modified[c] = w
}

func (m *Map) Flush() error {
	f, err := lockedfile.Edit(m.path)
	if err != nil {
		return fmt.Errorf("flush caching map: %s", err)
	}
	defer f.Close()
	if err := m.readFrom(f); err != nil {
		return fmt.Errorf("flush caching map %s: %s", m.path, err)
	}
	tmp := m.path + ".new"
	f2, err := os.Create(tmp)
	if err != nil {
		return fmt.Errorf("flush caching map: %s", err)
	}
	defer f2.Close()
	if err := m.writeTo(f2); err != nil {
		return fmt.Errorf("flush caching map %s: %s", tmp, err)
	}
	if err := f2.Close(); err != nil {
		return fmt.Errorf("flush caching map: %s", err)
	}
	if err := os.Rename(tmp, m.path); err != nil {
		return fmt.Errorf("flush caching map: %s", err)
	}
	return nil
}

func (m *Map) readFrom(f io.Reader) error {
	r, err := gzip.NewReader(f)
	switch err {
	case nil:
	case io.EOF:
		return nil
	default:
		return fmt.Errorf("read caching map: %s", err)
	}
	defer r.Close()
	if err := gob.NewDecoder(r).Decode(&m.cached); err != nil {
		return fmt.Errorf("read caching map: %s", err)
	}
	return nil
}

func (m *Map) writeTo(f io.Writer) error {
	merged := make(map[codes.WorkCode]interface{})
	for k, v := range m.cached {
		merged[k] = v
	}
	for k, v := range m.modified {
		merged[k] = v
	}
	w := gzip.NewWriter(f)
	defer w.Close()
	if err := gob.NewEncoder(w).Encode(merged); err != nil {
		return fmt.Errorf("write caching map: %s", err)
	}
	if err := w.Close(); err != nil {
		return fmt.Errorf("write caching map: %s", err)
	}
	return nil
}
