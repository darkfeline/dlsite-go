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

// Package codes contains the types for DLSite RJ codes and RE codes.
package codes

import "regexp"

// A WorkCode is a code for a DLSite work.
type WorkCode string

// An RJCode is an RJ code for a DLSite work.
// An RJCode is a valid WorkCode.
type RJCode string

var rjCodePat = regexp.MustCompile("RJ[0-9]+")

// ParseRJCode returns the first RJ code found in the string.  If no RJ code
// is found, returns an empty value.
func ParseRJCode(s string) RJCode {
	return RJCode(rjCodePat.FindString(s))
}

// An RECode is an RE code for a DLSite work.
// An RECode is a valid WorkCode.
type RECode string

var reCodePat = regexp.MustCompile("RE[0-9]+")

// ParseRECode returns the first RE code found in the string.  If no RE code
// is found, returns an empty value.
func ParseRECode(s string) RECode {
	return RECode(reCodePat.FindString(s))
}
