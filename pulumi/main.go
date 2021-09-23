package main

import (
	"github.com/mizzy/sock-shop/pulumi/lib/ecs"
	"github.com/mizzy/sock-shop/pulumi/lib/lb"
	"github.com/mizzy/sock-shop/pulumi/lib/service"
	"github.com/mizzy/sock-shop/pulumi/lib/vpc"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		err := vpc.NewVpc(ctx)
		err = lb.NewLb(ctx)
		err = ecs.NewEcs(ctx)
		err = service.NewServices(ctx)
		return err
	})
}
