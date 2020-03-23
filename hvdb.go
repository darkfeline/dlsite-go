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
	"go.felesatra.moe/dlsite/v2/internal/hvdb"
)

func fillWorkFromHVDB(w *Work, c codes.RJCode) error {
	hw, err := hvdb.FetchWork(c)
	if err != nil {
		return err
	}
	if hw.Code != "" {
		w.Code = codes.WorkCode(hw.Code)
	}
	if hw.Title != "" {
		w.Title = hw.Title
	}
	if hw.EnglishTitle != "" {
		w.EnglishTitle = hw.EnglishTitle
	}
	if hw.Circle != "" {
		w.Circle = hw.Circle
	}
	w.CVs = append(w.CVs, hw.CVs...)
	w.Tags = append(w.Tags, hw.Tags...)
	w.SFW = hw.SFW
	return nil
}
