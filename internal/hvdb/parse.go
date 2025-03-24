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
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"go.felesatra.moe/dlsite/v2/codes"
)

func parseWork(r io.Reader) (*Work, error) {
	d, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}
	return &Work{
		Code:         parseWorkCode(d),
		Title:        parseTitle(d),
		EnglishTitle: parseEnglishTitle(d),
		Circle:       parseCircle(d),
		Tags:         parseTags(d),
		CVs:          parseCVs(d),
		SFW:          parseSFW(d),
	}, nil
}

func parseWorkCode(d *goquery.Document) codes.WorkCode {
	s := d.Find("div.body-content").Find("h2").Text()
	return codes.ParseCode(s)
}

func parseTitle(d *goquery.Document) string {
	s := strings.TrimSpace(d.Find(`label[for="Name"]`).Next().Text())
	return strings.TrimSpace(s)
}

func parseEnglishTitle(d *goquery.Document) string {
	s := strings.TrimSpace(d.Find(`label[for="EngName"]`).Next().Text())
	return strings.TrimSpace(s)
}

func parseCircle(d *goquery.Document) string {
	s := d.Find(`label[for="Circle_Name"]`).Next().Text()
	return strings.TrimSpace(strings.Split(s, " / ")[0])
}

func parseTags(d *goquery.Document) []string {
	s, _ := d.Find("#TagsString").Attr("value")
	return strings.Split(s, ",")
}

func parseCVs(d *goquery.Document) []string {
	s, _ := d.Find("#CVsString").Attr("value")
	return strings.Split(s, ",")
}

func parseSFW(d *goquery.Document) bool {
	s, _ := d.Find("#SFW").Attr("checked")
	return strings.TrimSpace(s) == "checked"
}
