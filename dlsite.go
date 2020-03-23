// Copyright (C) 2020 Allen Li
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
	"go.felesatra.moe/dlsite/v2/codes"
	"go.felesatra.moe/dlsite/v2/internal/dlsite"
)

func fillWorkFromDLSite(w *Work, c codes.RJCode) error {
	dw, err := dlsite.FetchWork(c)
	if err != nil {
		return err
	}
	if dw.Code != "" {
		w.Code = codes.WorkCode(dw.Code)
	}
	if dw.Title != "" {
		w.Title = dw.Title
	}
	if dw.Circle != "" {
		w.Circle = dw.Circle
	}
	if dw.Series != "" {
		w.Series = dw.Series
	}
	if dw.Description != "" {
		w.Description = dw.Description
	}
	return nil
}
