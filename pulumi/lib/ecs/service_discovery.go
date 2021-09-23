package ecs

import (
	"github.com/mizzy/sock-shop/pulumi/lib/vpc"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/servicediscovery"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var PrivateDnsNamespace *servicediscovery.PrivateDnsNamespace

func newSeriveDiscovery(ctx *pulumi.Context) error {
	var err error
	PrivateDnsNamespace, err = servicediscovery.NewPrivateDnsNamespace(ctx, "local", &servicediscovery.PrivateDnsNamespaceArgs{
		Name: pulumi.String("local"),
		Vpc:  vpc.Vpc.ID(),
	})
	if err != nil {
		return err
	}

	return nil
}
