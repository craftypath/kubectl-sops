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
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/craftypath/sops-operator/pkg/apis/craftypath/v1alpha1"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type config struct {
	output      string
	dryRun      bool
	dryRunType  string
	kubectlArgs []string
	sopsArgs    []string
}

var (
	Version   = "dev"
	GitCommit = "HEAD"
	BuildDate = "unknown"

	outputRegexes = []*regexp.Regexp{
		regexp.MustCompile(`^-o(json|yaml)?$`),
		regexp.MustCompile(`^--output(?:=(json|yaml))?$`),
	}
	dryRunRegex            = regexp.MustCompile(`^--dry-run(?:=(none|server|client))?$`)
	allowedDryRunTypeRegex = regexp.MustCompile(`none|client|server`)

	sopsFileFormats = map[string]string{
		".yaml": "yaml",
		".yml":  "yaml",
		".json": "json",
		".ini":  "ini",
		".env":  "dotenv",
	}
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		// TODO usage
		fmt.Fprintln(os.Stderr, "no args specified")
		os.Exit(1)
	}
	if args[0] == "--version" {
		fmt.Printf("%s (commit=%s, date=%s)\n", Version, GitCommit, BuildDate)
		os.Exit(0)
	}

	config := parseArgs(args)
	err := execute(config)
	if err != nil {
		var exitCode int
		var msg string
		if err, ok := err.(*exec.ExitError); ok {
			exitCode = err.ExitCode()
			msg = string(err.Stderr)
		} else {
			exitCode = 1
			msg = err.Error()
		}
		fmt.Fprintln(os.Stderr, msg)
		os.Exit(exitCode)
	}
}

func execute(config *config) error {
	secret := &corev1.Secret{}
	if err := populateSecret(config, secret); err != nil {
		return err
	}

	sopsSecret, err := createSopsSecret(config, secret)
	if err != nil {
		return err
	}

	if err := applySopsSecret(config, sopsSecret); err != nil {
		return err
	}
	return nil
}

func applySopsSecret(config *config, sopsSecret *v1alpha1.SopsSecret) error {
	b, err := json.Marshal(sopsSecret)
	if err != nil {
		return err
	}

	args := []string{
		"create", "--save-config=false",
	}
	if config.dryRun {
		if config.dryRunType != "" {
			args = append(args, "--dry-run="+config.dryRunType)
		} else {
			args = append(args, "--dry-run")
		}
	}
	if config.output != "" {
		args = append(args, "--output", config.output)
	}
	args = append(args, "--filename", "-")

	command := exec.Command("kubectl", args...)
	command.Stdin = bytes.NewReader(b)
	output, err := command.Output()
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return nil
}

func createSopsSecret(config *config, secret *corev1.Secret) (*v1alpha1.SopsSecret, error) {
	sopsSecret := &v1alpha1.SopsSecret{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "craftypath.github.io/v1alpha1",
			Kind:       "SopsSecret",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      secret.Name,
			Namespace: secret.Namespace,
		},
		Spec: v1alpha1.SopsSecretSpec{
			Type:       secret.Type,
			StringData: make(map[string]string),
		},
	}

	for k, v := range secret.Data {
		format := determineFileFormat(k)
		args := []string{"--input-type", format, "--output-type", format}
		args = append(args, config.sopsArgs...)
		args = append(args, "-e", "/dev/stdin")
		command := exec.Command("sops", args...)
		command.Stdin = bytes.NewReader(v)
		output, err := command.Output()
		if err != nil {
			return nil, err
		}
		sopsSecret.Spec.StringData[k] = string(output)
	}
	return sopsSecret, nil
}

func populateSecret(config *config, secret *corev1.Secret) error {
	args := config.kubectlArgs
	args = append(args, "-o", "json", "--dry-run")
	output, err := exec.Command("kubectl", args...).Output()
	if err != nil {
		return err
	}
	return json.Unmarshal(output, secret)
}

func parseArgs(args []string) *config {
	config := &config{}

	for i := 0; i < len(args); i++ {
		arg := args[i]

		if arg == "--" {
			config.sopsArgs = args[i+1:]
			break
		}
		var outputMatch bool
		for _, regex := range outputRegexes {
			if submatch := regex.FindStringSubmatch(arg); submatch != nil {
				if submatch[1] != "" {
					config.output = submatch[1]
				} else {
					i++
					config.output = args[i]
				}
				outputMatch = true
				break
			}
		}
		if outputMatch {
			continue
		}

		if submatch := dryRunRegex.FindStringSubmatch(arg); submatch != nil {
			config.dryRun = true
			if submatch[1] != "" {
				config.dryRunType = submatch[1]
			} else {
				if len(args) > i+1 {
					nextArg := args[i+1]
					if allowedDryRunTypeRegex.MatchString(nextArg) {
						config.dryRunType = nextArg
						i++
					}
				}
			}
			continue
		}

		config.kubectlArgs = append(config.kubectlArgs, arg)
	}

	if config.dryRunType == "server" {
		// SopsSecret doesn't support server-side dry-run yet
		config.dryRunType = "client"
	}
	return config
}

func determineFileFormat(fileName string) string {
	ext := filepath.Ext(fileName)
	if format, exists := sopsFileFormats[ext]; exists {
		return format
	}
	return "binary"
}
