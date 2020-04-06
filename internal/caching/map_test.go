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

package caching

import (
	"encoding/gob"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"go.felesatra.moe/dlsite/v2/codes"
)

type Work struct {
	Title string
}

func init() {
	gob.Register((*Work)(nil))
}

func TestMap(t *testing.T) {
	t.Parallel()
	tempdir, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.RemoveAll(tempdir) })
	p := filepath.Join(tempdir, "map")
	m, err := Open(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { m.Close() })
	c := codes.WorkCode("RJ123")
	want := &Work{Title: "eyjafjalla"}
	m.Put(c, want)
	t.Run("get from modified", func(t *testing.T) {
		var got *Work
		m.Get(c, &got)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %+v; want %+v", got, want)
		}
	})
	m.Close()

	m, err = Open(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { m.Close() })
	t.Run("get from saved", func(t *testing.T) {
		var got *Work
		m.Get(c, &got)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %+v; want %+v", got, want)
		}
	})
	m.Close()
}
