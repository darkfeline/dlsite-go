/*
Package dlsite provides a Go API for accessing DLSite information.
*/
package dlsite

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/pkg/errors"
)

// RJCode is the type for RJ code strings.  Use Parse to get RJCode
// values from strings.
type RJCode string

func (c RJCode) workURL() string {
	return fmt.Sprintf("http://www.dlsite.com/maniax/work/=/product_id/%s.html", c)
}

func (c RJCode) announceURL() string {
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
	defer r.Body.Close()
	if err != nil {
		return nil, errors.Wrapf(err, "requesting %s", c)
	}
	w, err := parseWork(c, r.Body)
	if err != nil {
		return nil, errors.Wrap(err, "parse work")
	}
	return w, nil
}

// requestPage requests the page for a DLSite work.  If error is nil,
// make sure to defer a call to r.Body.Close().
func requestPage(c RJCode) (*http.Response, error) {
	r, err := http.Get(c.workURL())
	if err != nil {
		return nil, errors.Wrapf(err, "getting %s", c.workURL())
	}
	if r.StatusCode == 404 {
		return requestAnnouncePage(c)
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("GET %s: HTTP %d %s", c.workURL(), r.StatusCode, r.Status)
	}
	return r, nil
}

func requestAnnouncePage(c RJCode) (*http.Response, error) {
	r, err := http.Get(c.announceURL())
	if err != nil {
		return nil, errors.Wrapf(err, "getting %s", c.announceURL())
	}
	if r.StatusCode == 404 {
		return nil, errors.Errorf("Cannot find %s", c)
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("GET %s: HTTP %d %s", c.announceURL(), r.StatusCode, r.Status)
	}
	return r, nil
}
