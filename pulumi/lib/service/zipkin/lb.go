package zipkin

import (
	"github.com/mizzy/sock-shop/pulumi/lib/ecs"
	"github.com/mizzy/sock-shop/pulumi/lib/vpc"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/alb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var targetGroup *alb.TargetGroup

func newLb(ctx *pulumi.Context) error {
	var err error
	targetGroup, err = alb.NewTargetGroup(ctx, "zipkin", &alb.TargetGroupArgs{
		DeregistrationDelay:            pulumi.Int(300),
		LambdaMultiValueHeadersEnabled: pulumi.Bool(false),
		Name:                           pulumi.String("sock-Zipki-1I5ZBFCMMEA3L"),
		Port:                           pulumi.Int(9411),
		Protocol:                       pulumi.String("HTTP"),
		ProxyProtocolV2:                pulumi.Bool(false),
		SlowStart:                      pulumi.Int(0),
		TargetType:                     pulumi.String("ip"),
		VpcId:                          vpc.Vpc.ID(),
	})
	if err != nil {
		return err
	}

	listener, err := alb.NewListener(ctx, "zipkin", &alb.ListenerArgs{
		DefaultActions: alb.ListenerDefaultActionArray{
			&alb.ListenerDefaultActionArgs{
				TargetGroupArn: targetGroup.ID(),
				Type:           pulumi.String("forward"),
			},
		},
		LoadBalancerArn: ecs.Lb.Arn,
		Port:            pulumi.Int(9411),
		Protocol:        pulumi.String("HTTP"),
	})
	if err != nil {
		return err
	}

	_, err = alb.NewListenerRule(ctx, "zipkin", &alb.ListenerRuleArgs{
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
