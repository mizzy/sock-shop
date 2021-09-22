package ecs

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/cloudwatch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func newCloudWatchLogs(ctx *pulumi.Context) error {
	_, err := cloudwatch.NewLogGroup(ctx, "sock_shop", &cloudwatch.LogGroupArgs{
		Name:            pulumi.String("sock-shop"),
		RetentionInDays: pulumi.Int(7),
	})
	if err != nil {
		return err
	}

	return nil
}
