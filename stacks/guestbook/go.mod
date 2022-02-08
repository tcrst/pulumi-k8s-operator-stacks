module github.com/pulumi/examples/kubernetes-go-guestbook/components

go 1.14

replace "github.com/tcrst/pulumi-k8s-operator-stacks/stacks/addons" => /pulumi-k8s-operator-stacks/stacks/guestbook/addons
require (
	github.com/pulumi/pulumi-kubernetes/sdk/v3 v3.0.0
	github.com/pulumi/pulumi/sdk/v3 v3.24.1
)

