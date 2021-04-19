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

// Package dlsite provides a Go API for accessing DLSite information.
package dlsite

import (
	"fmt"
	"net/http"

	"go.felesatra.moe/dlsite/v2/codes"
)

func workURL(c codes.RJCode) string {
	return fmt.Sprintf("http://www.dlsite.com/maniax/work/=/product_id/%s.html", c)
}

func announceURL(c codes.RJCode) string {
	return fmt.Sprintf("http://www.dlsite.com/maniax/announce/=/product_id/%s.html", c)
}

// Work holds information about a DLSite work.  If information for a
// field is missing, it will be the zero value.
type Work struct {
	Code        codes.RJCode
	Title       string
	Circle      string
	Series      string
	Seiyuu      []string
	WorkFormats []string
	Description string
}

// FetchWork returns the Work for the codes.RJCode by parsing the
// corresponding DLSite page.
// Errors from the HTTP request or from parsing the HTML
// of the page are returned, but errors from parsing specific
// information from the page are not returned.
// Not all fields will be present for all works, like Series.
func FetchWork(c codes.RJCode) (*Work, error) {
	r, err := requestPage(c)
	if err != nil {
		return nil, fmt.Errorf("fetch from dlsite: %w", err)
	}
	defer r.Body.Close()
	w, err := parseWork(c, r.Body)
	if err != nil {
		return nil, fmt.Errorf("fetch from dlsite: %w", err)
	}
	return w, nil
}

// requestPage requests the page for a DLSite work.  If error is nil,
// make sure to defer a call to r.Body.Close().
func requestPage(c codes.RJCode) (*http.Response, error) {
	r, err := http.Get(workURL(c))
	if err != nil {
		return nil, fmt.Errorf("GET %s: %w", workURL(c), err)
	}
	if r.StatusCode == 404 {
		r.Body.Close()
		return requestAnnouncePage(c)
	}
	if r.StatusCode != 200 {
		r.Body.Close()
		return nil, fmt.Errorf("GET %s: HTTP %s", workURL(c), r.Status)
	}
	return r, nil
}

func requestAnnouncePage(c codes.RJCode) (*http.Response, error) {
	r, err := http.Get(announceURL(c))
	if err != nil {
		return nil, fmt.Errorf("GET %s: %w", workURL(c), err)
	}
	if r.StatusCode != 200 {
		r.Body.Close()
		return nil, fmt.Errorf("GET %s: HTTP %s", workURL(c), r.Status)
	}
	return r, nil
}
