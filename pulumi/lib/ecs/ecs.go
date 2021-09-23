package ecs

import (
	"github.com/mizzy/sock-shop/pulumi/lib/ecs/frontend"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func NewEcs(ctx *pulumi.Context) error {
	resources := []func(*pulumi.Context) error{
		newCloudWatchLogs,
		newDynamoDbTaskRole,
		newEcsTaskExecutionRole,
		newSeriveDiscovery,
		newCluster,
		newSecurityGroup,
	}

	for _, r := range resources {
		err := r(ctx)
		if err != nil {
			return err
		}

	}

	services := []func(*pulumi.Context, *iam.Role, *iam.Role) error{
		frontend.NewFrontEnd,
	}

	for _, s := range services {
		err := s(ctx, taskExecutionRole, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
