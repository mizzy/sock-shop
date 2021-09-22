package ecs

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ecs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func newCluster(ctx *pulumi.Context) error {
	_, err := ecs.NewCluster(ctx, "sock_shop", &ecs.ClusterArgs{
		Name: pulumi.String("Sock-Shop"),
	}, pulumi.Protect(true))
	if err != nil {
		return err
	}

	return nil
}
