// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/pulumi/examples/kubernetes-go-guestbook/components/addons"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		// Deploy k8s addons
		deployAddons(ctx)

		// Deploy GuestBookServices
		deployServices(ctx)

		return nil
	})
}

func deployAddons(ctx *pulumi.Context) error {
	err := addons.NewMetricServer(ctx)
	if err != nil {
		return err
	}
	return nil
}

func deployServices(ctx *pulumi.Context) error {
	// Redis leader Deployment + Service
	_, err := NewServiceDeployment(ctx, "redis-leader", &ServiceDeploymentArgs{
		Image: pulumi.String("redis"),
		Ports: pulumi.IntArray{pulumi.Int(6379)},
	})
	if err != nil {
		return err
	}

	// Redis replica Deployment + Service
	_, err = NewServiceDeployment(ctx, "redis-replica", &ServiceDeploymentArgs{
		Image: pulumi.String("pulumi/guestbook-redis-replica"),
		Ports: pulumi.IntArray{pulumi.Int(6379)},
	})
	if err != nil {
		return err
	}

	// Frontend Deployment + Service
	frontend, err := NewServiceDeployment(ctx, "frontend", &ServiceDeploymentArgs{
		Image:    pulumi.String("pulumi/guestbook-php-redis"),
		Ports:    pulumi.IntArray{pulumi.Int(80)},
		Replicas: pulumi.Int(2),
		Hpa:      true,
	})
	if err != nil {
		return err
	}

	ctx.Export("frontendIP", frontend.Service.Spec.ApplyT(
		func(spec *corev1.ServiceSpec) *string { return spec.ClusterIP }))

	return nil

}
