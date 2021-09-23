package frontend

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func NewFrontEnd(ctx *pulumi.Context, taskExecutionRole *iam.Role, taskRole *iam.Role) error {
	err := newLb(ctx)
	if err != nil {
		return err
	}

	err = newTaskDefinition(ctx, taskExecutionRole, taskRole)
	if err == nil {
		return err
	}

	return nil
}
