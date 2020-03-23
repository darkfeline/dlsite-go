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

// Package hvdb provides a Go API for accessing HVDB information.
package hvdb

import (
	"fmt"
	"net/http"

	"go.felesatra.moe/dlsite/v2/codes"
)

type Work struct {
	Code         codes.RJCode
	Title        string
	EnglishTitle string
	Circle       string
	CVs          []string
	Tags         []string
	SFW          bool
}

func FetchWork(c codes.RJCode) (*Work, error) {
	r, err := requestPage(c)
	if err != nil {
		return nil, fmt.Errorf("fetch from hvdb: %s", err)
	}
	defer r.Body.Close()
	w, err := parseWork(r.Body)
	if err != nil {
		return nil, fmt.Errorf("fetch from hvdb %s: %s", c, err)
	}
	return w, nil
}

// requestPage requests the page for a work.  If error is nil,
// make sure to defer a call to r.Body.Close().
func requestPage(c codes.RJCode) (*http.Response, error) {
	r, err := http.Get(hvdbURL(c))
	if err != nil {
		return nil, fmt.Errorf("request %s: %s", c, err)
	}
	if r.StatusCode != 200 {
		r.Body.Close()
		return nil, fmt.Errorf("request %s: %s", c, r.Status)
	}
	return r, nil
}

func hvdbURL(c codes.RJCode) string {
	return fmt.Sprintf("https://hvdb.me/Dashboard/WorkDetails/%s", c[2:])
}
