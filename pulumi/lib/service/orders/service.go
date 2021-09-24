package orders

import (
	e "github.com/mizzy/sock-shop/pulumi/lib/ecs"
	"github.com/mizzy/sock-shop/pulumi/lib/vpc"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ecs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func newService(ctx *pulumi.Context) error {
	_, err := ecs.NewService(ctx, "orders", &ecs.ServiceArgs{
		Cluster:                         e.Cluster.Arn,
		DeploymentMaximumPercent:        pulumi.Int(200),
		DeploymentMinimumHealthyPercent: pulumi.Int(100),
		DesiredCount:                    pulumi.Int(1),
		EnableEcsManagedTags:            pulumi.Bool(false),
		EnableExecuteCommand:            pulumi.Bool(false),
		Name:                            pulumi.String("sock-shop-OrdersService-vKEiumEDxDGY"),
		LaunchType:                      pulumi.String("FARGATE"),
		NetworkConfiguration: &ecs.ServiceNetworkConfigurationArgs{
			AssignPublicIp: pulumi.Bool(true),
			SecurityGroups: pulumi.StringArray{
				e.EcsSecurityGroup.ID(),
			},
			Subnets: pulumi.StringArray{
				vpc.PublicSubnet1.ID(),
				vpc.PublicSubnet2.ID(),
			},
		},
		SchedulingStrategy: pulumi.String("REPLICA"),
		ServiceRegistries: &ecs.ServiceServiceRegistriesArgs{
			RegistryArn: registry.Arn,
		},
		TaskDefinition:     taskDef.Arn,
		WaitForSteadyState: pulumi.Bool(false),
	})
	if err != nil {
		return err
	}

	return nil
}
