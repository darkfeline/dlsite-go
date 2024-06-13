// Copyright (C) 2019 Allen Li
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

package hvdb

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"go.felesatra.moe/dlsite/v2/codes"
)

func testdataPath(c codes.WorkCode) string {
	return filepath.Join("testdata", fmt.Sprintf("%s.html", c))
}

func TestParseWork(t *testing.T) {
	t.Parallel()
	c := codes.WorkCode("RJ222837")
	want := &Work{
		Code:         c,
		Title:        "籠の鳥",
		EnglishTitle: "Caged Bird",
		Circle:       "骨格ゼロ動物",
		Tags: []string{
			"brunette hair",
			"fellatio",
			"oral cumshot",
			"removed work",
			"shrine maiden",
			"swallowing",
			"tentacle",
		},
		CVs: []string{
			"Kazuna Sayaka",
			"計名さや香",
		},
	}
	testParseWork(t, c, want)
}

func TestParseWorkWithBools(t *testing.T) {
	t.Parallel()
	c := codes.WorkCode("RJ242172")
	want := &Work{
		Code:   c,
		Title:  "森の中で赤ずきんのお姉さんと～内緒の甘々癒され耳かき～",
		Circle: "甘音缶",
		Tags: []string{
			"binaural audio",
			"ear cleaning",
			"ear licking",
			"healing",
			"moe",
		},
		CVs: []string{"桃箱"},
		SFW: true,
	}
	testParseWork(t, c, want)
}

func testParseWork(t *testing.T, c codes.WorkCode, want *Work) {
	f, err := os.Open(testdataPath(c))
	if err != nil {
		t.Fatalf("Error opening test file: %s", err)
	}
	defer f.Close()
	got, err := parseWork(f)
	if err != nil {
		t.Fatalf("Error parsing work: %s", err)
	}
	if want.Code != got.Code {
		t.Errorf("Code = %#v; want %#v", got.Code, want.Code)
	}
	if want.Title != got.Title {
		t.Errorf("Title = %#v; want %#v", got.Title, want.Title)
	}
	if want.EnglishTitle != got.EnglishTitle {
		t.Errorf("EnglishTitle = %#v; want %#v", got.EnglishTitle, want.EnglishTitle)
	}
	if want.Circle != got.Circle {
		t.Errorf("Circle = %#v; want %#v", got.Circle, want.Circle)
	}
	if !reflect.DeepEqual(want.Tags, got.Tags) {
		t.Errorf("Tags = %#v; want %#v", got.Tags, want.Tags)
	}
	if !reflect.DeepEqual(want.CVs, got.CVs) {
		t.Errorf("CVs = %#v; want %#v", got.CVs, want.CVs)
	}
	if want.SFW != got.SFW {
		t.Errorf("SFW = %#v; want %#v", got.SFW, want.SFW)
	}
}
