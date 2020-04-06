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
	"bytes"
	"database/sql"
	"encoding/gob"
	"fmt"
	"log"

	"go.felesatra.moe/dlsite/v2/codes"

	_ "github.com/mattn/go-sqlite3"
)

type Map struct {
	path string
	db   *sql.DB
}

func Open(p string) (*Map, error) {
	m := &Map{
		path: p,
	}
	var err error
	m.db, err = sql.Open("sqlite3", p)
	if err != nil {
		return nil, fmt.Errorf("open cache: %s", err)
	}
	if err := m.ensureTable(); err != nil {
		return nil, fmt.Errorf("open cache: %s", err)
	}
	return m, nil
}

func (m *Map) ensureTable() error {
	_, err := m.db.Exec(`CREATE TABLE IF NOT EXISTS dlsite_works (
code TEXT PRIMARY KEY,
data BLOB
)`)
	return err
}

func (m *Map) Get(c codes.WorkCode, v interface{}) {
	r := m.db.QueryRow(`SELECT data FROM dlsite_works WHERE code = ?`, c)
	var d []byte
	if err := r.Scan(&d); err != nil {
		log.Printf("cache get %s: %s", c, err)
		return
	}
	if err := decode(d, v); err != nil {
		panic(err)
	}
}

func (m *Map) Put(c codes.WorkCode, v interface{}) {
	d, err := encode(v)
	if err != nil {
		panic(err)
	}
	_, err = m.db.Exec(`INSERT INTO dlsite_works (code, data) VALUES (?, ?)`, c, d)
	if err != nil {
		panic(err)
	}
}

func (m *Map) Close() error {
	if err := m.db.Close(); err != nil {
		return fmt.Errorf("close cache: %s", err)
	}
	return nil
}

func encode(v interface{}) ([]byte, error) {
	var b bytes.Buffer
	if err := gob.NewEncoder(&b).Encode(v); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func decode(b []byte, v interface{}) error {
	if err := gob.NewDecoder(bytes.NewReader(b)).Decode(v); err != nil {
		return err
	}
	return nil
}
