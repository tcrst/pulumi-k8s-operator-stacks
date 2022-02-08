package addons

import (
	apiregistrationv1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/apiregistration/v1"
	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	rbacv1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/rbac/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func NewMetricServer(ctx *pulumi.Context, opts ...pulumi.ResourceOption) error {
	addonDeploment := &AddonDeployment{}

	err := ctx.RegisterComponentResource(
		"go-guestbook:addons:MetricsServer",
		"metrics-server", addonDeploment, opts...)
	if err != nil {
		return err
	}

	addonDeploment.ServiceAccount, err = corev1.NewServiceAccount(ctx, "kube_systemMetrics_serverServiceAccount", &corev1.ServiceAccountArgs{
		ApiVersion: pulumi.String("v1"),
		Kind:       pulumi.String("ServiceAccount"),
		Metadata: &metav1.ObjectMetaArgs{
			Labels: pulumi.StringMap{
				"k8s-app": pulumi.String("metrics-server"),
			},
			Name:      pulumi.String("metrics-server"),
			Namespace: pulumi.String("kube-system"),
		},
	})
	if err != nil {
		return err
	}
	addonDeploment.ClusterRole, err = rbacv1.NewClusterRole(ctx, "system:aggregated_metrics_readerClusterRole", &rbacv1.ClusterRoleArgs{
		ApiVersion: pulumi.String("rbac.authorization.k8s.io/v1"),
		Kind:       pulumi.String("ClusterRole"),
		Metadata: &metav1.ObjectMetaArgs{
			Labels: pulumi.StringMap{
				"k8s-app": pulumi.String("metrics-server"),
				"rbac.authorization.k8s.io/aggregate-to-admin": pulumi.String("true"),
				"rbac.authorization.k8s.io/aggregate-to-edit":  pulumi.String("true"),
				"rbac.authorization.k8s.io/aggregate-to-view":  pulumi.String("true"),
			},
			Name: pulumi.String("system:aggregated-metrics-reader"),
		},
		Rules: rbacv1.PolicyRuleArray{
			&rbacv1.PolicyRuleArgs{
				ApiGroups: pulumi.StringArray{
					pulumi.String("metrics.k8s.io"),
				},
				Resources: pulumi.StringArray{
					pulumi.String("pods"),
					pulumi.String("nodes"),
				},
				Verbs: pulumi.StringArray{
					pulumi.String("get"),
					pulumi.String("list"),
					pulumi.String("watch"),
				},
			},
		},
	})
	if err != nil {
		return err
	}
	addonDeploment.ClusterRole, err = rbacv1.NewClusterRole(ctx, "system:metrics_serverClusterRole", &rbacv1.ClusterRoleArgs{
		ApiVersion: pulumi.String("rbac.authorization.k8s.io/v1"),
		Kind:       pulumi.String("ClusterRole"),
		Metadata: &metav1.ObjectMetaArgs{
			Labels: pulumi.StringMap{
				"k8s-app": pulumi.String("metrics-server"),
			},
			Name: pulumi.String("system:metrics-server"),
		},
		Rules: rbacv1.PolicyRuleArray{
			&rbacv1.PolicyRuleArgs{
				ApiGroups: pulumi.StringArray{
					pulumi.String(""),
				},
				Resources: pulumi.StringArray{
					pulumi.String("nodes/metrics"),
				},
				Verbs: pulumi.StringArray{
					pulumi.String("get"),
				},
			},
			&rbacv1.PolicyRuleArgs{
				ApiGroups: pulumi.StringArray{
					pulumi.String(""),
				},
				Resources: pulumi.StringArray{
					pulumi.String("pods"),
					pulumi.String("nodes"),
				},
				Verbs: pulumi.StringArray{
					pulumi.String("get"),
					pulumi.String("list"),
					pulumi.String("watch"),
				},
			},
		},
	})
	if err != nil {
		return err
	}
	addonDeploment.RoleBinding, err = rbacv1.NewRoleBinding(ctx, "kube_systemMetrics_server_auth_readerRoleBinding", &rbacv1.RoleBindingArgs{
		ApiVersion: pulumi.String("rbac.authorization.k8s.io/v1"),
		Kind:       pulumi.String("RoleBinding"),
		Metadata: &metav1.ObjectMetaArgs{
			Labels: pulumi.StringMap{
				"k8s-app": pulumi.String("metrics-server"),
			},
			Name:      pulumi.String("metrics-server-auth-reader"),
			Namespace: pulumi.String("kube-system"),
		},
		RoleRef: &rbacv1.RoleRefArgs{
			ApiGroup: pulumi.String("rbac.authorization.k8s.io"),
			Kind:     pulumi.String("Role"),
			Name:     pulumi.String("extension-apiserver-authentication-reader"),
		},
		Subjects: rbacv1.SubjectArray{
			&rbacv1.SubjectArgs{
				Kind:      pulumi.String("ServiceAccount"),
				Name:      pulumi.String("metrics-server"),
				Namespace: pulumi.String("kube-system"),
			},
		},
	})
	if err != nil {
		return err
	}
	addonDeploment.ClusterRoleBinding, err = rbacv1.NewClusterRoleBinding(ctx, "metrics_server:system:auth_delegatorClusterRoleBinding", &rbacv1.ClusterRoleBindingArgs{
		ApiVersion: pulumi.String("rbac.authorization.k8s.io/v1"),
		Kind:       pulumi.String("ClusterRoleBinding"),
		Metadata: &metav1.ObjectMetaArgs{
			Labels: pulumi.StringMap{
				"k8s-app": pulumi.String("metrics-server"),
			},
			Name: pulumi.String("metrics-server:system:auth-delegator"),
		},
		RoleRef: &rbacv1.RoleRefArgs{
			ApiGroup: pulumi.String("rbac.authorization.k8s.io"),
			Kind:     pulumi.String("ClusterRole"),
			Name:     pulumi.String("system:auth-delegator"),
		},
		Subjects: rbacv1.SubjectArray{
			&rbacv1.SubjectArgs{
				Kind:      pulumi.String("ServiceAccount"),
				Name:      pulumi.String("metrics-server"),
				Namespace: pulumi.String("kube-system"),
			},
		},
	})
	if err != nil {
		return err
	}
	addonDeploment.ClusterRoleBinding, err = rbacv1.NewClusterRoleBinding(ctx, "system:metrics_serverClusterRoleBinding", &rbacv1.ClusterRoleBindingArgs{
		ApiVersion: pulumi.String("rbac.authorization.k8s.io/v1"),
		Kind:       pulumi.String("ClusterRoleBinding"),
		Metadata: &metav1.ObjectMetaArgs{
			Labels: pulumi.StringMap{
				"k8s-app": pulumi.String("metrics-server"),
			},
			Name: pulumi.String("system:metrics-server"),
		},
		RoleRef: &rbacv1.RoleRefArgs{
			ApiGroup: pulumi.String("rbac.authorization.k8s.io"),
			Kind:     pulumi.String("ClusterRole"),
			Name:     pulumi.String("system:metrics-server"),
		},
		Subjects: rbacv1.SubjectArray{
			&rbacv1.SubjectArgs{
				Kind:      pulumi.String("ServiceAccount"),
				Name:      pulumi.String("metrics-server"),
				Namespace: pulumi.String("kube-system"),
			},
		},
	})
	if err != nil {
		return err
	}
	addonDeploment.Service, err = corev1.NewService(ctx, "kube_systemMetrics_serverService", &corev1.ServiceArgs{
		ApiVersion: pulumi.String("v1"),
		Kind:       pulumi.String("Service"),
		Metadata: &metav1.ObjectMetaArgs{
			Labels: pulumi.StringMap{
				"k8s-app": pulumi.String("metrics-server"),
			},
			Name:      pulumi.String("metrics-server"),
			Namespace: pulumi.String("kube-system"),
		},
		Spec: &corev1.ServiceSpecArgs{
			Ports: corev1.ServicePortArray{
				&corev1.ServicePortArgs{
					Name:       pulumi.String("https"),
					Port:       pulumi.Int(443),
					Protocol:   pulumi.String("TCP"),
					TargetPort: pulumi.Any("https"),
				},
			},
			Selector: pulumi.StringMap{
				"k8s-app": pulumi.String("metrics-server"),
			},
		},
	})
	if err != nil {
		return err
	}
	addonDeploment.Deployment, err = appsv1.NewDeployment(ctx, "kube_systemMetrics_serverDeployment", &appsv1.DeploymentArgs{
		ApiVersion: pulumi.String("apps/v1"),
		Kind:       pulumi.String("Deployment"),
		Metadata: &metav1.ObjectMetaArgs{
			Labels: pulumi.StringMap{
				"k8s-app": pulumi.String("metrics-server"),
			},
			Name:      pulumi.String("metrics-server"),
			Namespace: pulumi.String("kube-system"),
		},
		Spec: &appsv1.DeploymentSpecArgs{
			Selector: &metav1.LabelSelectorArgs{
				MatchLabels: pulumi.StringMap{
					"k8s-app": pulumi.String("metrics-server"),
				},
			},
			Strategy: &appsv1.DeploymentStrategyArgs{
				RollingUpdate: &appsv1.RollingUpdateDeploymentArgs{
					MaxUnavailable: pulumi.Any(0),
				},
			},
			Template: &corev1.PodTemplateSpecArgs{
				Metadata: &metav1.ObjectMetaArgs{
					Labels: pulumi.StringMap{
						"k8s-app": pulumi.String("metrics-server"),
					},
				},
				Spec: &corev1.PodSpecArgs{
					Containers: corev1.ContainerArray{
						&corev1.ContainerArgs{
							Args: pulumi.StringArray{
								pulumi.String("--cert-dir=/tmp"),
								pulumi.String("--secure-port=4443"),
								pulumi.String("--kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname"),
								pulumi.String("--kubelet-use-node-status-port"),
								pulumi.String("--metric-resolution=15s"),
								pulumi.String("--kubelet-insecure-tls"),
							},
							Image:           pulumi.String("k8s.gcr.io/metrics-server/metrics-server:v0.6.0"),
							ImagePullPolicy: pulumi.String("IfNotPresent"),
							LivenessProbe: &corev1.ProbeArgs{
								FailureThreshold: pulumi.Int(3),
								HttpGet: &corev1.HTTPGetActionArgs{
									Path:   pulumi.String("/livez"),
									Port:   pulumi.Any("https"),
									Scheme: pulumi.String("HTTPS"),
								},
								PeriodSeconds: pulumi.Int(10),
							},
							Name: pulumi.String("metrics-server"),
							Ports: corev1.ContainerPortArray{
								&corev1.ContainerPortArgs{
									ContainerPort: pulumi.Int(4443),
									Name:          pulumi.String("https"),
									Protocol:      pulumi.String("TCP"),
								},
							},
							ReadinessProbe: &corev1.ProbeArgs{
								FailureThreshold: pulumi.Int(3),
								HttpGet: &corev1.HTTPGetActionArgs{
									Path:   pulumi.String("/readyz"),
									Port:   pulumi.Any("https"),
									Scheme: pulumi.String("HTTPS"),
								},
								InitialDelaySeconds: pulumi.Int(20),
								PeriodSeconds:       pulumi.Int(10),
							},
							Resources: &corev1.ResourceRequirementsArgs{
								Requests: pulumi.StringMap{
									"cpu":    pulumi.String("100m"),
									"memory": pulumi.String("200Mi"),
								},
							},
							SecurityContext: &corev1.SecurityContextArgs{
								AllowPrivilegeEscalation: pulumi.Bool(false),
								ReadOnlyRootFilesystem:   pulumi.Bool(true),
								RunAsNonRoot:             pulumi.Bool(true),
								RunAsUser:                pulumi.Int(1000),
							},
							VolumeMounts: corev1.VolumeMountArray{
								&corev1.VolumeMountArgs{
									MountPath: pulumi.String("/tmp"),
									Name:      pulumi.String("tmp-dir"),
								},
							},
						},
					},
					NodeSelector: pulumi.StringMap{
						"kubernetes.io/os": pulumi.String("linux"),
					},
					PriorityClassName:  pulumi.String("system-cluster-critical"),
					ServiceAccountName: pulumi.String("metrics-server"),
					Volumes: corev1.VolumeArray{
						&corev1.VolumeArgs{
							EmptyDir: nil,
							Name:     pulumi.String("tmp-dir"),
						},
					},
				},
			},
		},
	})
	if err != nil {
		return err
	}
	addonDeploment.ApiRegistration, err = apiregistrationv1.NewAPIService(ctx, "v1beta1_metrics_k8s_ioAPIService", &apiregistrationv1.APIServiceArgs{
		ApiVersion: pulumi.String("apiregistration.k8s.io/v1"),
		Kind:       pulumi.String("APIService"),
		Metadata: &metav1.ObjectMetaArgs{
			Labels: pulumi.StringMap{
				"k8s-app": pulumi.String("metrics-server"),
			},
			Name: pulumi.String("v1beta1.metrics.k8s.io"),
		},
		Spec: &apiregistrationv1.APIServiceSpecArgs{
			Group:                 pulumi.String("metrics.k8s.io"),
			GroupPriorityMinimum:  pulumi.Int(100),
			InsecureSkipTLSVerify: pulumi.Bool(true),
			Service: &apiregistrationv1.ServiceReferenceArgs{
				Name:      pulumi.String("metrics-server"),
				Namespace: pulumi.String("kube-system"),
			},
			Version:         pulumi.String("v1beta1"),
			VersionPriority: pulumi.Int(100),
		},
	})
	if err != nil {
		return err
	}
	return nil
}
