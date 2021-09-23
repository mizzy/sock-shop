package service

import (
	"github.com/mizzy/sock-shop/pulumi/lib/service/frontend"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func NewServices(ctx *pulumi.Context) error {
	frontend.NewFrontEnd(ctx)

	return nil
}
