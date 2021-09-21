package ecs

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

func NewEcs(ctx *pulumi.Context) error {
	err := cloudWatchLogs(ctx)
	if err != nil {
		return err
	}

	err = dynamoDbTaskRole(ctx)
	if err != nil {
		return err
	}

	return nil
}
