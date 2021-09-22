package ecs

import (
	"github.com/mizzy/sock-shop/pulumi/lib/vpc"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func newSecurityGroup(ctx *pulumi.Context) error {
	_, err := ec2.NewSecurityGroup(ctx, "ecs", &ec2.SecurityGroupArgs{
		Description:         pulumi.String("ECS Allowed Ports"),
		Name:                pulumi.String("sock-shop-EcsSecurityGroup-1JN0GGW02EK2G"),
		RevokeRulesOnDelete: pulumi.Bool(false),
		VpcId:               vpc.Vpc.ID(),
		Tags: pulumi.StringMap{
			"Name": pulumi.String("ecs"),
		},
	})
	if err != nil {
		return err
	}

	return nil
}
