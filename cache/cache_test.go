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

package cache

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"go.felesatra.moe/dlsite"
)

func TestGetMissingWork(t *testing.T) {
	t.Parallel()
	d, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Fatalf("Error making temp dir: %s", err)
	}
	defer os.RemoveAll(d)
	p := filepath.Join(d, "tmp.db")
	c, err := Open(p)
	if err != nil {
		t.Fatalf("Error opening cache: %s", err)
	}
	defer c.Close()
	_, err = c.Get(dlsite.RJCode("RJ1234"))
	if err == nil {
		t.Errorf("Got nil error when getting missing key")
	}
}

func TestPutAndGetWork(t *testing.T) {
	t.Parallel()
	d, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Fatalf("Error making temp dir: %s", err)
	}
	defer os.RemoveAll(d)
	p := filepath.Join(d, "tmp.db")
	c, err := Open(p)
	if err != nil {
		t.Fatalf("Error opening cache: %s", err)
	}
	defer c.Close()
	r := dlsite.RJCode("RJ1234")
	w := &dlsite.Work{
		RJCode: r,
		Name:   "foobar",
	}
	err = c.Put(w)
	if err != nil {
		t.Fatalf("Error putting work: %s", err)
	}
	got, err := c.Get(r)
	if err != nil {
		t.Fatalf("Error getting work: %s", err)
	}
	if !reflect.DeepEqual(w, got) {
		t.Errorf("Expected %#v, got %#v", w, got)
	}
}

func TestKeys(t *testing.T) {
	t.Parallel()
	d, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(d)
	p := filepath.Join(d, "tmp.db")
	c, err := Open(p)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	w := &dlsite.Work{
		RJCode: "RJ1234",
		Name:   "foobar",
	}
	err = c.Put(w)
	if err != nil {
		t.Fatal(err)
	}
	w = &dlsite.Work{
		RJCode: "RJ1235",
		Name:   "foobar",
	}
	err = c.Put(w)
	if err != nil {
		t.Fatal(err)
	}
	ks, err := c.Keys()
	if err != nil {
		t.Fatal(err)
	}
	want := []dlsite.RJCode{"RJ1234", "RJ1235"}
	if !reflect.DeepEqual(ks, want) {
		t.Errorf("Got %#v; want %#v", ks, want)
	}
}
