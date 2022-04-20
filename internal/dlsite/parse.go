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

package dlsite

import (
	"fmt"
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"go.felesatra.moe/dlsite/v2/codes"
)

func parseWork(c codes.RJCode, r io.Reader) (*Work, error) {
	d, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, fmt.Errorf("cannot parse document: %w", err)
	}
	w := &Work{
		Code:        c,
		Title:       parseTitle(d),
		Circle:      parseCircle(d),
		Series:      parseSeries(d),
		Seiyuu:      parseSeiyuu(d),
		WorkFormats: parseWorkFormats(d),
		Description: parseDescription(d),
	}
	return w, nil
}

func parseTitle(d *goquery.Document) string {
	return strings.TrimSpace(d.Find("#work_name").Text())
}

func parseCircle(d *goquery.Document) string {
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

func parseSeiyuu(d *goquery.Document) []string {
	fields := strings.Split(
		d.Find("#work_outline").
			Find(`th:contains("声優")`).Next().Text(),
		"/")
	for i := range fields {
		fields[i] = strings.TrimSpace(fields[i])
	}
	if len(fields) == 1 && fields[0] == "" {
		return nil
	}
	return fields
}

func parseWorkFormats(d *goquery.Document) []string {
	var f []string
	spans := d.Find("div.work_genre").
		Slice(1, 2).
		Find("span")
	for i := 0; i < spans.Length(); i++ {
		n := spans.Get(i)
		f = append(f, n.FirstChild.Data)
	}
	return f
}

func parseDescription(d *goquery.Document) string {
	return strings.TrimSpace(d.
		Find("#main_inner").
		Find(`div[itemprop="description"]`).Text())
}
