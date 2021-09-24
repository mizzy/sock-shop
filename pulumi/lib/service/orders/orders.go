package orders

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

func NewOrders(ctx *pulumi.Context) error {
	resources := []func(ctx2 *pulumi.Context) error{
		newDynamoDB,
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
