package main

import (
	"github.com/mizzy/sock-shop/pulumi/lib/ecs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		err := ecs.NewEcs(ctx)
		return err
	})
}
