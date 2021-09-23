package frontend

import (
	"github.com/mizzy/sock-shop/pulumi/lib/lb"
	"github.com/mizzy/sock-shop/pulumi/lib/vpc"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/alb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var targetGroup *alb.TargetGroup

func newLb(ctx *pulumi.Context) error {
	lb, err := alb.NewLoadBalancer(ctx, "lb", &alb.LoadBalancerArgs{
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

	targetGroup, err = alb.NewTargetGroup(ctx, "frontend", &alb.TargetGroupArgs{
		DeregistrationDelay:            pulumi.Int(300),
		LambdaMultiValueHeadersEnabled: pulumi.Bool(false),
		Name:                           pulumi.String("sock-Front-1QDQY0UJQC5EN"),
		Port:                           pulumi.Int(8079),
		Protocol:                       pulumi.String("HTTP"),
		ProxyProtocolV2:                pulumi.Bool(false),
		SlowStart:                      pulumi.Int(0),
		TargetType:                     pulumi.String("ip"),
		VpcId:                          vpc.Vpc.ID(),
	})
	if err != nil {
		return err
	}

	listener, err := alb.NewListener(ctx, "frontend", &alb.ListenerArgs{
		DefaultActions: alb.ListenerDefaultActionArray{
			&alb.ListenerDefaultActionArgs{
				TargetGroupArn: targetGroup.ID(),
				Type:           pulumi.String("forward"),
			},
		},
		LoadBalancerArn: lb.Arn,
		Port:            pulumi.Int(80),
		Protocol:        pulumi.String("HTTP"),
	})
	if err != nil {
		return err
	}

	_, err = alb.NewListenerRule(ctx, "frontend", &alb.ListenerRuleArgs{
		Actions: &alb.ListenerRuleActionArray{
			&alb.ListenerRuleActionArgs{
				TargetGroupArn: targetGroup.Arn,
				Type:           pulumi.String("forward"),
			},
		},
		Conditions: &alb.ListenerRuleConditionArray{
			&alb.ListenerRuleConditionArgs{
				PathPattern: &alb.ListenerRuleConditionPathPatternArgs{
					Values: pulumi.StringArray{
						pulumi.String("*"),
					},
				},
			},
		},
		ListenerArn: listener.Arn,
		Priority:    pulumi.Int(1),
	})
	if err != nil {
		return err
	}

	return nil
}
