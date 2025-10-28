module github.com/gigvault/operator

go 1.23

require (
	github.com/gigvault/shared v0.0.0
	k8s.io/apimachinery v0.29.0
	k8s.io/client-go v0.29.0
	sigs.k8s.io/controller-runtime v0.17.0
)

replace github.com/gigvault/shared => ../shared
