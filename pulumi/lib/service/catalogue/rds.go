package catalogue

import (
	"github.com/mizzy/sock-shop/pulumi/lib/ecs"
	"github.com/mizzy/sock-shop/pulumi/lib/vpc"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func newRds(ctx *pulumi.Context) error {
	sg, err := ec2.NewSecurityGroup(ctx, "db_ecs", &ec2.SecurityGroupArgs{
		Description:         pulumi.String("Open database for access"),
		Name:                pulumi.String("sock-shop-DBEC2SecurityGroup-9O8Q86URMQK3"),
		RevokeRulesOnDelete: pulumi.Bool(false),
		Tags: pulumi.StringMap{
			"Name": pulumi.String("db-ecs"),
		},
		VpcId: vpc.Vpc.ID(),
	})
	if err != nil {
		return err
	}

	_, err = ec2.NewSecurityGroupRule(ctx, "db_ecs_allow_from_sg", &ec2.SecurityGroupRuleArgs{
		FromPort:              pulumi.Int(3306),
		Protocol:              pulumi.String("tcp"),
		SecurityGroupId:       sg.ID(),
		ToPort:                pulumi.Int(3306),
		Type:                  pulumi.String("ingress"),
		SourceSecurityGroupId: ecs.EcsSecurityGroup.ID(),
	})
	if err != nil {
		return err
	}

	return nil
}
