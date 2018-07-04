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
)

// ReferenceNumber contains a reference number as used on the orange
// payment slip (Einzahlungsschein mit Referenznummer, ESR) in Switzerland.
// It implements fmt.Stringer interface as well as the interface
// structref.Printer defined in this package.
type ReferenceNumber struct {
	number string
}

// NewReferenceNumber creates a filled-in validated CreditorReference,
// or returns an error if the supplied ref string does not represent
// a valid creditor reference.
func NewReferenceNumber(ref string) (*ReferenceNumber, error) {
	stripped := stripSpaces(ref)
	if err := validateReferenceNumber(stripped); err != nil {
		return nil, err
	}
	return &ReferenceNumber{number: stripped}, nil
}

// NewReferenceNumberOrDie either returns a filled-in validated reference
// number, or aborts the program on error.
func NewReferenceNumberOrDie(ref string) *ReferenceNumber {
	nr, err := NewReferenceNumber(ref)
	if err != nil {
		panic(err)
	}
	return nr
}

// DigitalFormat returns the reference number digital format: no spaces.
func (ref *ReferenceNumber) DigitalFormat() string {
	return ref.number
}

// PrintFormat returns the reference number in space-separated groups.
func (ref *ReferenceNumber) PrintFormat() string {
	return addSpaces(ref.DigitalFormat(), 2, 5)
}

// String returns a plain string representation of a CreditorReference.
func (ref *ReferenceNumber) String() string {
	return ref.DigitalFormat()
}

// validateReferenceNumber checks that a reference number is valid for ESR.
// The input string must not contain any spaces.
//
// For details on the theory behind this validation see the PhD Thesis by Damm:
// Total anti-symmetrische Quasigruppen, Philipps-Univärsität Marburg, 2004.
func validateReferenceNumber(s string) error {
	matched, err := regexp.MatchString(`^[0-9]+$`, s)
	if err != nil {
		return fmt.Errorf("regexp error: %v", err)
	}
	if !matched {
		return fmt.Errorf("Illegal character in reference number: %v", s)
	}
	if len(s) < 2 {
		return fmt.Errorf("Too few digits in reference number: %v", s)
	}
	state := 0
	alg := []int{0, 9, 4, 6, 8, 2, 7, 1, 3, 5}
	for i := 0; i < len(s)-1; i++ {
		state = alg[(state+int(s[i]-'0'))%10]
	}
	checksum := 10 - state
	if checksum != int(s[len(s)-1]-'0') {
		return fmt.Errorf("Checksum mismatch in reference number: %v", s)
	}
	return nil
}
