package queuemaster

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

func NewQueueMaster(ctx *pulumi.Context) error {
	resources := []func(ctx2 *pulumi.Context) error{
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
