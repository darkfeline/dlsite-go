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

	"go.felesatra.moe/go2/errors"

	"go.felesatra.moe/dlsite"
)

// Work holds information about a work.  If information for a field is
// missing, it will be the zero value.
type Work struct {
	RJCode       dlsite.RJCode
	Title        string
	EnglishTitle string
	Circle       string
	Tags         []string
	CVs          []string
	SFW          bool
}

// Fetch fetches the Work by parsing the corresponding HVDB page.
// Errors from the HTTP request or from parsing the HTML of the page
// are returned, but errors from parsing specific information from the
// page are not returned.  Not all fields will be present for all
// works.
func Fetch(c dlsite.RJCode) (*Work, error) {
	r, err := requestPage(c)
	if err != nil {
		return nil, errors.Wrapf(err, "requesting %s", c)
	}
	defer r.Body.Close()
	w, err := parseWork(r.Body)
	if err != nil {
		return nil, errors.Wrap(err, "parse work")
	}
	return w, nil
}

// requestPage requests the page for a DLSite work.  If error is nil,
// make sure to defer a call to r.Body.Close().
func requestPage(c dlsite.RJCode) (*http.Response, error) {
	r, err := http.Get(hvdbURL(c))
	if err != nil {
		return nil, errors.Wrapf(err, "getting %s", hvdbURL(c))
	}
	if r.StatusCode == 404 {
		r.Body.Close()
		return nil, fmt.Errorf("Cannot find %s", c)
	}
	if r.StatusCode != 200 {
		r.Body.Close()
		return nil, fmt.Errorf("GET %s: HTTP %d %s", hvdbURL(c), r.StatusCode, r.Status)
	}
	return r, nil
}

func hvdbURL(c dlsite.RJCode) string {
	return fmt.Sprintf("https://hvdb.me/Dashboard/WorkDetails/%s", c[2:])
}
