/*
Copyright 2019 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package extract

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestExtract(t *testing.T) {
	// The file to scan
	paths := []string{"testdata/sample_file.go"}

	// The function we care about
	functions := []string{"extract.PrintToScreen"}

	// The directory where the sample translation file is in
	output := "testdata/"

	expected := map[string]interface{}{
		"Hint: This is not a URL, come on.":         "",
		"Holy cow I'm in a loop!":                   "Something else",
		"This is a variable with a string assigned": "",
		"This was a choice: %s":                     "Something",
		"Wow another string: %s":                    "",
	}

	err := TranslatableStrings(paths, functions, output)

	if err != nil {
		t.Fatalf("Error translating strings: %v", err)
	}

	var got map[string]interface{}
	f, err := ioutil.ReadFile("testdata/test.json")
	if err != nil {
		t.Fatalf("Reading json file: %s", err)
	}

	err = json.Unmarshal(f, &got)
	if err != nil {
		t.Fatalf("Error unmarshalling json: %v", err)
	}

	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("Translation JSON not equal: expected %v, got %v", expected, got)
	}

}
