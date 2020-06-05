# kubectl-sops

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

A `kubectl` plugin for creating `SopsSecret` resources.

See https://github.com/craftypath/sops-operator.

The plugin automatically encrypts data using [Mozilla SOPS](https://github.com/mozilla/sops) and wraps them into a `SopsSecret`.
The interface is the same as that of `kubectl create secret`.

## Installation

Download a release for your platform and add it to the `PATH`.
A distribution via [Krew](https://krew.sigs.k8s.io/) is planned.

## Examples

### From literal values

```console
kubectl sops create secret generic test-secret --from-literal foo=foo_secret --from-literal bar=bar_secret
```

### From file

```console
kubectl sops create secret generic test-secret --from-file test.yaml
```

### From file printing resulting YAML without applying it

```console
kubectl sops create secret generic test-secret --from-literal foo.yaml="bar: barvalue" --dry-run -o yaml
```

### With additional parameters for SOPS

* Useful if no `.sops.yaml` is used
* Args after the `--` delimiter are passed to SOPS

```console
kubectl sops create secret generic test-secret --from-file test.yaml -- \
    --kms arn:aws:kms:eu-central-1:123456789012:key/ffad06af-a6cc-43e5-ad61-51db75d17c77
```
