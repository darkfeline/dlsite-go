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
		RJCode:      c,
		Name:        parseName(d),
		Circle:      parseCircle(d),
		Series:      parseSeries(d),
		Description: parseDescription(d),
	}
	return w, nil
}

func parseName(d *goquery.Document) string {
	return strings.TrimSpace(d.Find("#work_name").Find("a").Text())
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

func parseDescription(d *goquery.Document) string {
	return strings.TrimSpace(d.
		Find("#main_inner").
		Find(`div[itemprop="description"]`).Text())
}
