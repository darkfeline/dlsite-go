package dlsite

import (
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

func parseWork(c RJCode, r io.Reader) (*Work, error) {
	d, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot parse document")
	}
	w := &Work{
		RJCode:      c,
		Name:        parseName(d),
		Maker:       parseMaker(d),
		Series:      parseSeries(d),
		Description: parseDescription(d),
		TrackList:   parseTrackList(d),
	}
	return w, nil
}

func parseName(d *goquery.Document) string {
	return strings.TrimSpace(d.Find("#work_name").Find("a").Text())
}

func parseMaker(d *goquery.Document) string {
	return strings.TrimSpace(d.
		Find("#work_maker").
		Find(`span[class="maker_name"]`).
		Find("a").Text())
}

func parseSeries(d *goquery.Document) string {
	return strings.TrimSpace(d.
		Find("#work_outline").
		Find(`th:contains("シリーズ名")`).Next().Text())
}

func parseDescription(d *goquery.Document) string {
	return strings.TrimSpace(d.
		Find("#main_inner").
		Find(`div[itemprop="description"]`).Text())
}

func parseTrackList(d *goquery.Document) []Track {
	var tl []Track
	s := d.Find("#work_parts").Find("ol.work_tracklist_list").ChildrenFiltered("li")
	s.Each(func(_ int, s *goquery.Selection) {
		s.Find("p.track_name").Find("span.track_number").Remove()
		t := Track{
			Name: s.Find("p.track_name").Text(),
			Text: s.Find("p.track_text").Text(),
		}
		tl = append(tl, t)
	})
	return tl
}
