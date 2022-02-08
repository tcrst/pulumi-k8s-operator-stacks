package addons

import (
	apiregistrationv1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/apiregistration/v1"
	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	rbacv1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/rbac/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AddonDeployment struct {
	pulumi.ResourceState
	Deployment         *appsv1.Deployment
	Service            *corev1.Service
	ServiceAccount     *corev1.ServiceAccount
	ClusterRole        *rbacv1.ClusterRole
	RoleBinding        *rbacv1.RoleBinding
	ClusterRoleBinding *rbacv1.ClusterRoleBinding
	ApiRegistration    *apiregistrationv1.APIService
}
