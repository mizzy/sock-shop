package vpc

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func newRoutes(ctx *pulumi.Context) error {
	internetGateway, err := ec2.NewInternetGateway(ctx, "internet_gateway", &ec2.InternetGatewayArgs{
		VpcId: Vpc.ID(),
	})
	if err != nil {
		return err
	}

	routeTable, err := ec2.NewRouteTable(ctx, "route_via_igw", &ec2.RouteTableArgs{
		VpcId: Vpc.ID(),
	})
	if err != nil {
		return err
	}

	_, err = ec2.NewRoute(ctx, "default", &ec2.RouteArgs{
		DestinationCidrBlock: pulumi.String("0.0.0.0/0"),
		GatewayId:            internetGateway.ID(),
		RouteTableId:         routeTable.ID(),
	})
	if err != nil {
		return err
	}

	_, err = ec2.NewRouteTableAssociation(ctx, "public_subnet_1_via_igw", &ec2.RouteTableAssociationArgs{
		RouteTableId: routeTable.ID(),
		SubnetId:     PublicSubnet1.ID(),
	})
	if err != nil {
		return err
	}

	_, err = ec2.NewRouteTableAssociation(ctx, "public_subnet_2_via_igw", &ec2.RouteTableAssociationArgs{
		RouteTableId: routeTable.ID(),
		SubnetId:     PublicSubnet2.ID(),
	})
	if err != nil {
		return err
	}

	return nil
}
