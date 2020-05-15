module github.com/craftypath/kubectl-sops

go 1.14

require (
	github.com/craftypath/sops-operator v0.3.0
	github.com/golangci/golangci-lint v1.25.0
	github.com/goreleaser/goreleaser v0.132.1
	github.com/magefile/mage v1.9.0
	golang.org/x/tools v0.0.0-20200422022333-3d57cf2e726e
	k8s.io/api v0.17.4
	k8s.io/apimachinery v0.17.4
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.3.2+incompatible // Required by OLM
	k8s.io/client-go => k8s.io/client-go v0.17.4 // Required by prometheus-operator
)
