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

// Package structref contains data types and validation functions for
// structured references: IBAN numbers, creditor references as described
// in ISO 11649, and Swiss ESR numbers.
package structref

// Printer defines methods for outputting structured identifiers.
type Printer interface {
	// A format suitable for digital distribution.
	// This typically means a string with no spaces.
	DigitalFormat() string

	// A format suitable for print.
	// This typically means a string with space-separated groups of
	// characters for more easy reading and human verification.
	PrintFormat() string
}
