package ecs

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

func NewEcs(ctx *pulumi.Context) error {
	resources := []func(ctx *pulumi.Context) error{
		cloudWatchLogs,
		dynamoDbTaskRole,
	}

	for _, r := range resources {
		err := r(ctx)
		if err != nil {
			return err
		}

	}

	return nil
}
