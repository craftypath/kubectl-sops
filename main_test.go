// Copyright The kubectl-sops Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"testing"
)

func Test_parseArgs(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{
			name: "-oyaml",
			args: []string{
				"create", "secret", "-oyaml",
			},
		},
		{
			name: "-ojson",
			args: []string{
				"create", "secret", "-ojson",
			},
		},
		{
			name: "-o yaml",
			args: []string{
				"create", "secret", "-o", "yaml",
			},
		},
		{
			name: "-o json",
			args: []string{
				"create", "secret", "-o", "json",
			},
		},
		{
			name: "--output=yaml",
			args: []string{
				"create", "secret", "--output=yaml",
			},
		},
		{
			name: "--output=json",
			args: []string{
				"create", "secret", "--output=json",
			},
		},
		{
			name: "--output yaml",
			args: []string{
				"create", "secret", "--output", "yaml",
			},
		},
		{
			name: "--output json",
			args: []string{
				"create", "secret", "--output", "json",
			},
		},
		{
			name: "--dry-run",
			args: []string{
				"create", "secret", "--dry-run",
			},
		},
		{
			name: "--dry-run=client",
			args: []string{
				"create", "secret", "--dry-run=client",
			},
		},
		{
			name: "--dry-run client",
			args: []string{
				"create", "secret", "--dry-run", "client",
			},
		},
		{
			name: "--dry-run -o yaml",
			args: []string{
				"create", "secret", "--dry-run", "-o", "yaml",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO write assertions
			parseArgs(tt.args)
		})
	}
}
