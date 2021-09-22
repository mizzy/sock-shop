package vpc

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func NewVpc(ctx *pulumi.Context) error {
	_, err := ec2.NewVpc(ctx, "sock-shop", &ec2.VpcArgs{
		AssignGeneratedIpv6CidrBlock: pulumi.Bool(false),
		CidrBlock:                    pulumi.String("172.31.0.0/16"),
		EnableDnsSupport:             pulumi.Bool(true),
		EnableDnsHostnames:           pulumi.Bool(true),
		InstanceTenancy:              pulumi.String("default"),
		Tags: pulumi.StringMap{
			"Name": pulumi.String("sock-shop"),
		},
	})
	if err != nil {
		return err
	}

	return nil
}
