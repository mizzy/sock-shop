package lb

import (
	"github.com/mizzy/sock-shop/pulumi/lib/vpc"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var ElbSecurityGroup *ec2.SecurityGroup

func newSecurityGroup(ctx *pulumi.Context) error {
	var err error
	ElbSecurityGroup, err = ec2.NewSecurityGroup(ctx, "elb_allowed_ports", &ec2.SecurityGroupArgs{
		Description:         pulumi.String("ELB Allowed Ports"),
		Name:                pulumi.String("sock-shop-ElbSecurityGroup-IK7UY8L0AXDX"),
		RevokeRulesOnDelete: pulumi.Bool(false),
		VpcId:               vpc.Vpc.ID(),
		Tags: pulumi.StringMap{
			"Name": pulumi.String("elb-allowed-ports"),
		},
	})
	if err != nil {
		return err
	}

	_, err = ec2.NewSecurityGroupRule(ctx, "elb_allowed_port_http", &ec2.SecurityGroupRuleArgs{
		CidrBlocks: pulumi.StringArray{
			pulumi.String("0.0.0.0/0"),
		},
		FromPort:        pulumi.Int(80),
		Protocol:        pulumi.String("tcp"),
		SecurityGroupId: ElbSecurityGroup.ID(),
		ToPort:          pulumi.Int(80),
		Type:            pulumi.String("ingress"),
	})
	if err != nil {
		return err
	}

	_, err = ec2.NewSecurityGroupRule(ctx, "elb_allowed_port_zipkin", &ec2.SecurityGroupRuleArgs{
		CidrBlocks: pulumi.StringArray{
			pulumi.String("0.0.0.0/0"),
		},
		FromPort:        pulumi.Int(9411),
		Protocol:        pulumi.String("tcp"),
		SecurityGroupId: ElbSecurityGroup.ID(),
		ToPort:          pulumi.Int(9411),
		Type:            pulumi.String("ingress"),
	})
	if err != nil {
		return err
	}

	_, err = ec2.NewSecurityGroupRule(ctx, "elb_allow_to_all", &ec2.SecurityGroupRuleArgs{
		CidrBlocks: pulumi.StringArray{
			pulumi.String("0.0.0.0/0"),
		},
		FromPort:        pulumi.Int(0),
		Protocol:        pulumi.String("-1"),
		SecurityGroupId: ElbSecurityGroup.ID(),
		ToPort:          pulumi.Int(0),
		Type:            pulumi.String("egress"),
	}, pulumi.Protect(true))
	if err != nil {
		return err
	}

	return nil
}
