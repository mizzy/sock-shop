package zipkin

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

func NewZipkin(ctx *pulumi.Context) error {
	resources := []func(ctx2 *pulumi.Context) error{
		newLb,
		newTaskDefinition,
		newRegistry,
		newService,
	}

	for _, r := range resources {
		err := r(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
