package frontend

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

func NewFrontEnd(ctx *pulumi.Context) error {
	resources := []func(ctx *pulumi.Context) error{
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
