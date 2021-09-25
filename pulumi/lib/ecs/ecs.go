package ecs

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func NewEcs(ctx *pulumi.Context) error {
	resources := []func(*pulumi.Context) error{
		newCloudWatchLogs,
		newEcsTaskExecutionRole,
		newSeriveDiscovery,
		newCluster,
		newSecurityGroup,
		newLb,
	}

	for _, r := range resources {
		err := r(ctx)
		if err != nil {
			return err
		}

	}

	return nil
}
