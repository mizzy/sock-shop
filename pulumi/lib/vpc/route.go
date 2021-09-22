package vpc

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func newRoutes(ctx *pulumi.Context) error {
	_, err := ec2.NewInternetGateway(ctx, "internet_gateway", &ec2.InternetGatewayArgs{
		VpcId: Vpc.ID(),
	})
	if err != nil {
		return err
	}

	return nil
}
