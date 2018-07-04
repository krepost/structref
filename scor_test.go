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

import "fmt"

func ExampleCreditorReference() {
	ref, err := NewCreditorReference("RF18 5390 0754 7034")
	fmt.Println(ref.DigitalFormat())
	fmt.Println(ref.PrintFormat())
	fmt.Println(err)
	// Output:
	// RF18539007547034
	// RF18 5390 0754 7034
	// <nil>
}

func ExamplePaddedCreditorReference() {
	ref, err := NewPaddedCreditorReference("RF18 5390 0754 7034")
	fmt.Println(ref.DigitalFormat())
	fmt.Println(ref.PrintFormat())
	fmt.Println(err)
	// Output:
	// RF18000000000539007547034
	// RF18 0000 0000 0539 0075 4703 4
	// <nil>
}

func ExampleErroneousCreditorReference() {
	ref, err := NewCreditorReference("RF6801")
	fmt.Println(ref)
	fmt.Println(err)
	// Output:
	// <nil>
	// Invalid checksum in creditor reference: RF6801
}
