package lb

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

func NewLb(ctx *pulumi.Context) error {
	err := newSecurityGroup(ctx)
	if err != nil {
		return err
	}

	return nil
}
