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
	"testing"
)

func TestEmptyInput(t *testing.T) {
	if ret := addSpaces("", 3, 17); "" != ret {
		t.Errorf("addSpaces() failed on empty input, returned %#v", ret)
	}
	if ret := stripSpaces(""); "" != ret {
		t.Errorf("stripSpaces() failed on empty input, returned %#v", ret)
	}
}

func TestRemovalOfSimpleWhiteSpace(t *testing.T) {
	input := "abc def 123 456"
	expected := "abcdef123456"
	if ret := stripSpaces(input); expected != ret {
		t.Errorf("stripSpaces(%#v): expected %#v, got %#v", input, expected, ret)
	}
}

func TestRemovalOfArbitraryWhiteSpace(t *testing.T) {
	input := "abc\t\ndef\f\t123\u2000456"
	expected := "abcdef123456"
	if ret := stripSpaces(input); expected != ret {
		t.Errorf("stripSpaces(%#v): expected %#v, got %#v", input, expected, ret)
	}
}

func TestRemovalOfArbitraryWhiteSpaceWithNonASCIICharacters(t *testing.T) {
	input := "αβγ\t\nδεζ\f\tабв\u2000где"
	expected := "αβγδεζабвгде"
	if ret := stripSpaces(input); expected != ret {
		t.Errorf("stripSpaces(%#v): expected %#v, got %#v", input, expected, ret)
	}
}

func TestGrouping(t *testing.T) {
	input := "abcdef123456"
	expected := "ab cde f12 345 6"
	if ret := addSpaces(input, 2, 3); expected != ret {
		t.Errorf("addSpaces(%#v): expected %#v, got %#v", input, expected, ret)
	}
}

func TestGroupingWithNonASCIICharacters(t *testing.T) {
	input := "αβγδεζабвгде"
	expected := "αβ γδε ζаб вгд е"
	if ret := addSpaces(input, 2, 3); expected != ret {
		t.Errorf("addSpaces(%#v): expected %#v, got %#v", input, expected, ret)
	}
}

func TestBigGroups(t *testing.T) {
	if ret := addSpaces("ab", 3, 17); "ab" != ret {
		t.Errorf("addSpaces() failed on big groups, returned %#v", ret)
	}
	if ret := addSpaces("abcd", 3, 17); "abc d" != ret {
		t.Errorf("addSpaces() failed on big groups, returned %#v", ret)
	}
}

func TestNegativeGroupSizes(t *testing.T) {
	if ret := addSpaces("abcdefgh", -2, 2); "abcdefgh" != ret {
		t.Errorf("addSpaces() failed on negative head length, returned %#v", ret)
	}
	if ret := addSpaces("abcdefgh", 2, -2); "abcdefgh" != ret {
		t.Errorf("addSpaces() failed on negative tail length, returned %#v", ret)
	}
}
