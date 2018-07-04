// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the “License”);
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an “AS IS” BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package structref

import (
	"strings"
	"unicode"
)

// addSpaces returns a print format version of a structured reference.
// Spaces are inserted so that the returned string contains a first
// group of head runes and then groups of tail runes each.
func addSpaces(normalized string, head, tail int) string {
	if head <= 0 || tail <= 0 {
		return normalized
	}
	chunk := head
	result := []rune{}
	for _, r := range normalized {
		if chunk == 0 {
			result = append(result, ' ')
			chunk = tail
		}
		result = append(result, r)
		chunk = chunk - 1
	}
	return string(result)
}

// stripSpaces removes all white-space from the input string.
func stripSpaces(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)
}
