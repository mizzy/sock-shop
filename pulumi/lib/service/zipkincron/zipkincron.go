package zipkincron

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

func NewZipkinCron(ctx *pulumi.Context) error {
	resources := []func(ctx2 *pulumi.Context) error{
		newTaskDefinition,
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
