package ecs

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ecs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var Cluster *ecs.Cluster

func newCluster(ctx *pulumi.Context) error {
	var err error
	Cluster, err = ecs.NewCluster(ctx, "sock_shop", &ecs.ClusterArgs{
		Name: pulumi.String("Sock-Shop"),
	})
	if err != nil {
		return err
	}

	return nil
}
