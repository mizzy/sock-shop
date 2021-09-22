package vpc

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var Vpc *ec2.Vpc
var PublicSubnet1 *ec2.Subnet
var PublicSubnet2 *ec2.Subnet

func NewVpc(ctx *pulumi.Context) error {
	resources := []func(ctx *pulumi.Context) error{
		newVpc,
		newSubnets,
		newRoutes,
	}

	for _, r := range resources {
		err := r(ctx)
		if err != nil {
			return err
		}

	}

	return nil

}

func newVpc(ctx *pulumi.Context) error {
	var err error
	Vpc, err = ec2.NewVpc(ctx, "sock_shop", &ec2.VpcArgs{
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

	dhcpOptions, err := ec2.NewVpcDhcpOptions(ctx, "local", &ec2.VpcDhcpOptionsArgs{
		DomainName: pulumi.String("local"),
		DomainNameServers: pulumi.StringArray{
			pulumi.String("AmazonProvidedDNS"),
		},
	})
	if err != nil {
		return err
	}

	_, err = ec2.NewVpcDhcpOptionsAssociation(ctx, "local", &ec2.VpcDhcpOptionsAssociationArgs{
		DhcpOptionsId: dhcpOptions.ID(),
		VpcId:         Vpc.ID(),
	})
	if err != nil {
		return err
	}

	return nil
}

func newSubnets(ctx *pulumi.Context) error {
	var err error
	PublicSubnet1, err = ec2.NewSubnet(ctx, "public_subnet_1", &ec2.SubnetArgs{
		AssignIpv6AddressOnCreation: pulumi.Bool(false),
		CidrBlock:                   pulumi.String("172.31.0.0/24"),
		MapPublicIpOnLaunch:         pulumi.Bool(false),
		AvailabilityZone:            pulumi.String("ap-northeast-1a"),
		Tags: pulumi.StringMap{
			"Name": pulumi.String("public-subnet-1"),
		},
		VpcId: Vpc.ID(),
	})
	if err != nil {
		return err
	}

	PublicSubnet2, err = ec2.NewSubnet(ctx, "public_subnet_2", &ec2.SubnetArgs{
		AssignIpv6AddressOnCreation: pulumi.Bool(false),
		CidrBlock:                   pulumi.String("172.31.1.0/24"),
		MapPublicIpOnLaunch:         pulumi.Bool(false),
		AvailabilityZone:            pulumi.String("ap-northeast-1c"),
		Tags: pulumi.StringMap{
			"Name": pulumi.String("public-subnet-2"),
		},
		VpcId: Vpc.ID(),
	})
	if err != nil {
		return err
	}

	return nil
}
