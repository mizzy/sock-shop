package catalogue

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

func NewCatalogue(ctx *pulumi.Context) error {
	resources := []func(ctx2 *pulumi.Context) error{
		newRds,
		newTaskDefinition,
		//newRegistry,
		//newService,
	}

	for _, r := range resources {
		err := r(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
