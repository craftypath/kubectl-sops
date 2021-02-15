module github.com/craftypath/kubectl-sops

go 1.15

require (
	github.com/craftypath/sops-operator v0.7.0
	github.com/golangci/golangci-lint v1.27.0
	github.com/goreleaser/goreleaser v0.138.0
	github.com/magefile/mage v1.10.0
	github.com/stretchr/testify v1.6.1
	golang.org/x/tools v0.0.0-20200608174601-1b747fd94509
	k8s.io/api v0.18.2
	k8s.io/apimachinery v0.18.2
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.3.2+incompatible // Required by OLM
	k8s.io/client-go => k8s.io/client-go v0.18.2 // Required by prometheus-operator
)
