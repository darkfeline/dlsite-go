// Copyright (C) 2024 Allen Li
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

package codes

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseCode(t *testing.T) {
	t.Parallel()
	cases := []struct {
		input string
		code  WorkCode
	}{
		{"blahRJ8642blah", "RJ8642"},
		{"blahRE8642blah", "RE8642"},
		{"blahBJ8642blah", "BJ8642"},
	}
	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			t.Parallel()
			got := ParseCode(c.input)
			if diff := cmp.Diff(c.code, got); diff != "" {
				t.Errorf("ParseCode() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestWorkType(t *testing.T) {
	t.Parallel()
	cases := []struct {
		code WorkCode
		wt   string
	}{
		{"RJ1234", "RJ"},
		{"RE1234", "RE"},
		{"BJ1234", "BJ"},
	}
	for _, c := range cases {
		t.Run(string(c.code), func(t *testing.T) {
			t.Parallel()
			got := c.code.WorkType()
			if diff := cmp.Diff(c.wt, got); diff != "" {
				t.Errorf("WorkType() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
