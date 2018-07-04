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
	"fmt"
	"regexp"
	"strings"
)

// CreditorReference contains a creditor reference as described in ISO 11649.
// It implements fmt.Stringer interface as well as the interface
// structref.Printer defined in this package.
type CreditorReference struct {
	// The reference root: everything except the “RF” prefix and check digits.
	root string
	// The check digits of the reference: the two digits following the
	// “RF” prefix.
	checkDigits string
}

// NewCreditorReference creates a filled-in validated CreditorReference,
// or returns an error if the supplied ref string does not represent
// a valid creditor reference.
func NewCreditorReference(ref string) (*CreditorReference, error) {
	normalized := strings.ToUpper(stripSpaces(ref))
	if err := validateCreditorReference(normalized); err != nil {
		return nil, err
	}
	creditor := CreditorReference{
		root:        normalized[4:],
		checkDigits: normalized[2:4],
	}
	return &creditor, nil
}

// NewCreditorReferenceOrDie either returns a filled-in validated creditor
// reference, or aborts the program on error.
func NewCreditorReferenceOrDie(ref string) *CreditorReference {
	creditor, err := NewCreditorReference(ref)
	if err != nil {
		panic(err)
	}
	return creditor
}

// NewPaddedCreditorReference is like NewCreditorReference but guarantees
// that the returned CreditorReference is exactly 25 characters by padding
// the root with zeros.
func NewPaddedCreditorReference(ref string) (*CreditorReference, error) {
	stripped := stripSpaces(ref)
	if len(stripped) < 4 || len(stripped) > 25 {
		return NewCreditorReference(ref) // Will return an error to caller.
	}
	root := stripped[4:]
	padding := strings.Repeat("0", 21-len(root))
	return NewCreditorReference(ref[0:4] + padding + root)
}

// NewPaddedCreditorReferenceOrDie either returns a filled-in validated
// creditor reference, or aborts the program on error.
func NewPaddedCreditorReferenceOrDie(ref string) *CreditorReference {
	creditor, err := NewPaddedCreditorReference(ref)
	if err != nil {
		panic(err)
	}
	return creditor
}

// DigitalFormat returns the creditor reference in digital format:
// the “RF” prefix followed by the two check digits and then the root.
func (ref *CreditorReference) DigitalFormat() string {
	return "RF" + ref.checkDigits + ref.root
}

// PrintFormat returns the creditor reference in space-separated groups.
func (ref *CreditorReference) PrintFormat() string {
	return addSpaces(ref.DigitalFormat(), 4, 4)
}

// String returns a plain string representation of a CreditorReference.
func (ref *CreditorReference) String() string {
	return ref.DigitalFormat()
}

// validateCreditorReference checks that a creditor reference is valid
// according to ISO 11649.
func validateCreditorReference(s string) error {
	matched, err := regexp.MatchString(`^[A-Z0-9]+$`, s)
	if err != nil {
		return fmt.Errorf("regexp error: %v", err)
	}
	if !matched {
		return fmt.Errorf("Illegal character in creditor reference: %v", s)
	}
	if len(s) < 4 {
		return fmt.Errorf("No prefix and checksum in creditor reference: %v", s)
	}
	if len(s) > 25 {
		return fmt.Errorf("Too many characters in creditor reference: %v", s)
	}
	if s[0:2] != "RF" {
		return fmt.Errorf("Incorrect prefix in creditor reference: %v", s)
	}
	modulus := 0
	for i := 0; i < len(s); i++ {
		ch := s[(i+4)%len(s)]
		if ch >= 'A' && ch <= 'Z' {
			modulus = (100*modulus + int((ch-'A')+10)) % 97
		} else if ch >= '0' && ch <= '9' {
			modulus = (10*modulus + int(ch-'0')) % 97
		}
	}
	if modulus != 1 {
		return fmt.Errorf("Invalid checksum in creditor reference: %v", s)
	}
	return nil
}
