package frontend

import (
	"github.com/mizzy/sock-shop/pulumi/lib/lb"
	"github.com/mizzy/sock-shop/pulumi/lib/vpc"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/alb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func newLb(ctx *pulumi.Context) error {
	_, err := alb.NewLoadBalancer(ctx, "lb", &alb.LoadBalancerArgs{
		DropInvalidHeaderFields:      pulumi.Bool(false),
		EnableCrossZoneLoadBalancing: pulumi.Bool(false),
		EnableDeletionProtection:     pulumi.Bool(false),
		EnableHttp2:                  pulumi.Bool(true),
		IdleTimeout:                  pulumi.Int(30),
		LoadBalancerType:             pulumi.String("application"),
		Name:                         pulumi.String("sockshop"),
		Subnets: pulumi.StringArray{
			vpc.PublicSubnet1.ID(),
			vpc.PublicSubnet2.ID(),
		},
		SecurityGroups: pulumi.StringArray{lb.ElbSecurityGroup.ID()},
	})
	if err != nil {
		return err
	}

	return nil
}
