package zipkin

import (
	"github.com/mizzy/sock-shop/pulumi/lib/vpc"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/alb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var targetGroup *alb.TargetGroup

func newLb(ctx *pulumi.Context) error {
	_, err := alb.NewTargetGroup(ctx, "zipkin", &alb.TargetGroupArgs{
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

	return nil
}
