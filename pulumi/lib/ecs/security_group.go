package ecs

import (
	"github.com/mizzy/sock-shop/pulumi/lib/lb"
	"github.com/mizzy/sock-shop/pulumi/lib/vpc"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func newSecurityGroup(ctx *pulumi.Context) error {
	sg, err := ec2.NewSecurityGroup(ctx, "ecs", &ec2.SecurityGroupArgs{
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

	_, err = ec2.NewSecurityGroupRule(ctx, "ecs_allow_ssh_from_all", &ec2.SecurityGroupRuleArgs{
		CidrBlocks: pulumi.StringArray{
			pulumi.String("0.0.0.0/0"),
		},
		FromPort:        pulumi.Int(22),
		Protocol:        pulumi.String("tcp"),
		SecurityGroupId: sg.ID(),
		ToPort:          pulumi.Int(22),
		Type:            pulumi.String("ingress"),
	})
	if err != nil {
		return err
	}

	_, err = ec2.NewSecurityGroupRule(ctx, "ecs_allow_all_from_elb", &ec2.SecurityGroupRuleArgs{
		FromPort:              pulumi.Int(0),
		Protocol:              pulumi.String("-1"),
		SecurityGroupId:       sg.ID(),
		SourceSecurityGroupId: lb.ElbSecurityGroup.ID(),
		ToPort:                pulumi.Int(0),
		Type:                  pulumi.String("ingress"),
	})
	if err != nil {
		return err
	}

	return nil
}
