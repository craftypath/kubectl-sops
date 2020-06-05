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

	"github.com/stretchr/testify/assert"
)

func Test_parseArgs(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want *config
	}{
		{
			name: "-oyaml",
			args: []string{
				"create", "secret", "-oyaml",
			},
			want: &config{
				output:      "yaml",
				dryRun:      false,
				dryRunType:  "",
				kubectlArgs: []string{"create", "secret"},
			},
		},
		{
			name: "-ojson",
			args: []string{
				"create", "secret", "-ojson",
			},
			want: &config{
				output:      "json",
				dryRun:      false,
				dryRunType:  "",
				kubectlArgs: []string{"create", "secret"},
			},
		},
		{
			name: "-o yaml",
			args: []string{
				"create", "secret", "-o", "yaml",
			},
			want: &config{
				output:      "yaml",
				dryRun:      false,
				dryRunType:  "",
				kubectlArgs: []string{"create", "secret"},
			},
		},
		{
			name: "-o json",
			args: []string{
				"create", "secret", "-o", "json",
			},
			want: &config{
				output:      "json",
				dryRun:      false,
				dryRunType:  "",
				kubectlArgs: []string{"create", "secret"},
			},
		},
		{
			name: "--output=yaml",
			args: []string{
				"create", "secret", "--output=yaml",
			},
			want: &config{
				output:      "yaml",
				dryRun:      false,
				dryRunType:  "",
				kubectlArgs: []string{"create", "secret"},
			},
		},
		{
			name: "--output=json",
			args: []string{
				"create", "secret", "--output=json",
			},
			want: &config{
				output:      "json",
				dryRun:      false,
				dryRunType:  "",
				kubectlArgs: []string{"create", "secret"},
			},
		},
		{
			name: "--output yaml",
			args: []string{
				"create", "secret", "--output", "yaml",
			},
			want: &config{
				output:      "yaml",
				dryRun:      false,
				dryRunType:  "",
				kubectlArgs: []string{"create", "secret"},
			},
		},
		{
			name: "--output json",
			args: []string{
				"create", "secret", "--output", "json",
			},
			want: &config{
				output:      "json",
				dryRun:      false,
				dryRunType:  "",
				kubectlArgs: []string{"create", "secret"},
			},
		},
		{
			name: "--dry-run",
			args: []string{
				"create", "secret", "--dry-run",
			},
			want: &config{
				output:      "",
				dryRun:      true,
				dryRunType:  "",
				kubectlArgs: []string{"create", "secret"},
			},
		},
		{
			name: "--dry-run=client",
			args: []string{
				"create", "secret", "--dry-run=client",
			},
			want: &config{
				output:      "",
				dryRun:      true,
				dryRunType:  "client",
				kubectlArgs: []string{"create", "secret"},
			},
		},
		{
			name: "--dry-run client",
			args: []string{
				"create", "secret", "--dry-run", "client",
			},
			want: &config{
				output:      "",
				dryRun:      true,
				dryRunType:  "client",
				kubectlArgs: []string{"create", "secret"},
			},
		},
		{
			name: "--dry-run server",
			args: []string{
				"create", "secret", "--dry-run", "server",
			},
			want: &config{
				output:      "",
				dryRun:      true,
				dryRunType:  "client", // server is not supported -> use client
				kubectlArgs: []string{"create", "secret"},
			},
		},
		{
			name: "--dry-run -o yaml",
			args: []string{
				"create", "secret", "--dry-run", "-o", "yaml",
			},
			want: &config{
				output:      "yaml",
				dryRun:      true,
				dryRunType:  "",
				kubectlArgs: []string{"create", "secret"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseArgs(tt.args)
			assert.Equal(t, tt.want, got)
		})
	}
}
