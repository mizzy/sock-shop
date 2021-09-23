package ecs

import (
	"github.com/mizzy/sock-shop/pulumi/lib/ecs/frontend"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func NewEcs(ctx *pulumi.Context) error {
	resources := []func(ctx *pulumi.Context) error{
		newCloudWatchLogs,
		newDynamoDbTaskRole,
		newEcsTaskExecutionRole,
		newSeriveDiscovery,
		newCluster,
		newSecurityGroup,
		frontend.NewFrontEnd,
	}

	for _, r := range resources {
		err := r(ctx)
		if err != nil {
			return err
		}

	}

	return nil
}
