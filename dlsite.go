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

/*
Package dlsite provides a Go API for accessing DLSite information.
*/
package dlsite

import (
	"fmt"
	"net/http"
	"regexp"
)

// RJCode is the type for RJ code strings.  Use Parse to get RJCode
// values from strings.
type RJCode string

func workURL(c RJCode) string {
	return fmt.Sprintf("http://www.dlsite.com/maniax/work/=/product_id/%s.html", c)
}

func announceURL(c RJCode) string {
	return fmt.Sprintf("http://www.dlsite.com/maniax/announce/=/product_id/%s.html", c)
}

// Work holds information about a DLSite work.  If information for a
// field is missing, it will be the zero value.
type Work struct {
	RJCode      RJCode
	Name        string
	Maker       string
	Series      string
	Description string
	TrackList   []Track
}

// Track holds information about a single track.
type Track struct {
	Name string
	Text string
}

var rjCodePat = regexp.MustCompile("RJ[0-9]+")

// Parse returns the first RJ code found in the string.  If no RJ code
// is found, Parse returns an empty value.
func Parse(s string) RJCode {
	return RJCode(rjCodePat.FindString(s))
}

// Fetch returns the Work for the RJCode by parsing the corresponding
// DLSite page.  Errors from the HTTP request or from parsing the HTML
// of the page are returned, but errors from parsing specific
// information from the page are not returned.  Not all fields will be
// present for all works, like Series or TrackList.
func Fetch(c RJCode) (*Work, error) {
	r, err := requestPage(c)
	if err != nil {
		return nil, fmt.Errorf("requesting %s: %w", c, err)
	}
	defer r.Body.Close()
	w, err := parseWork(c, r.Body)
	if err != nil {
		return nil, fmt.Errorf("parse work: %w", err)
	}
	return w, nil
}

// requestPage requests the page for a DLSite work.  If error is nil,
// make sure to defer a call to r.Body.Close().
func requestPage(c RJCode) (*http.Response, error) {
	r, err := http.Get(workURL(c))
	if err != nil {
		return nil, fmt.Errorf("getting %s: %w", workURL(c), err)
	}
	if r.StatusCode == 404 {
		r.Body.Close()
		return requestAnnouncePage(c)
	}
	if r.StatusCode != 200 {
		r.Body.Close()
		return nil, fmt.Errorf("GET %s: HTTP %d %s", workURL(c), r.StatusCode, r.Status)
	}
	return r, nil
}

func requestAnnouncePage(c RJCode) (*http.Response, error) {
	r, err := http.Get(announceURL(c))
	if err != nil {
		return nil, fmt.Errorf("getting %s: %w", announceURL(c), err)
	}
	if r.StatusCode == 404 {
		r.Body.Close()
		return nil, fmt.Errorf("Cannot find %s", c)
	}
	if r.StatusCode != 200 {
		r.Body.Close()
		return nil, fmt.Errorf("GET %s: HTTP %d %s", announceURL(c), r.StatusCode, r.Status)
	}
	return r, nil
}
