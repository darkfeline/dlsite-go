// Copyright (C) 2021 Allen Li
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
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRemoveDupes(t *testing.T) {
	t.Parallel()
	cases := []struct {
		input []string
		want  []string
	}{
		{
			input: []string{"a", "b", "a"},
			want:  []string{"a", "b"},
		},
		{
			input: []string{"b", "a", "a"},
			want:  []string{"a", "b"},
		},
		{
			input: []string{"a", "a", "b", "b"},
			want:  []string{"a", "b"},
		},
		{
			input: []string{"b", "c", "a"},
			want:  []string{"a", "b", "c"},
		},
	}
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("%v", c.input), func(t *testing.T) {
			t.Parallel()
			got := removeDupes(c.input)
			if diff := cmp.Diff(c.want, got); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
