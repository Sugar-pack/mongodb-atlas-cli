// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resources

import (
	"strings"

	"k8s.io/apimachinery/pkg/util/validation"
)

func AtlasNameToKubernetesName() map[string]string {
	return map[string]string{
		" ": "-",
		".": "dot",
		"@": "at",
		"(": "left-parenthesis",
		")": "right-parenthesis",
		"&": "and",
		"+": "plus",
		":": "colon",
		",": "comma",
		"'": "single-quote",
	}
}

// NormalizeAtlasName normalizes the name to be used as a resource name in Kubernetes
// Run fuzzing test if you want to change this function!
func NormalizeAtlasName(name string, dictionary map[string]string) string {
	for k, v := range dictionary {
		name = strings.ReplaceAll(name, k, v)
	}

	restrictionForFirstAndLast := []string{"-", "_"}
	dictionaryForFirstAndLast := map[string]string{
		restrictionForFirstAndLast[0]: "dash",
		restrictionForFirstAndLast[1]: "underscore",
	}
	if len(name) > 0 {
		for _, v := range restrictionForFirstAndLast {
			if strings.HasPrefix(name, v) {
				name = dictionaryForFirstAndLast[v] + name[1:]
			}
			if strings.HasSuffix(name, v) {
				name = name[:len(name)-1] + dictionaryForFirstAndLast[v]
			}
		}
	}

	if len(name) > validation.DNS1123LabelMaxLength {
		name = name[:validation.DNS1123LabelMaxLength]
	}

	if len(name) > 0 && name[len(name)-1] == '-' {
		name = name[:len(name)-1] + "d"
	}

	return strings.ToLower(name)
}
