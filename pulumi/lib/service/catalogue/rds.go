package catalogue

import (
	"github.com/mizzy/sock-shop/pulumi/lib/ecs"
	"github.com/mizzy/sock-shop/pulumi/lib/vpc"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func newRds(ctx *pulumi.Context) error {
	sg, err := ec2.NewSecurityGroup(ctx, "db_ecs", &ec2.SecurityGroupArgs{
		Description:         pulumi.String("Open database for access"),
		Name:                pulumi.String("sock-shop-DBEC2SecurityGroup-9O8Q86URMQK3"),
		RevokeRulesOnDelete: pulumi.Bool(false),
		Tags: pulumi.StringMap{
			"Name": pulumi.String("db-ecs"),
		},
		VpcId: vpc.Vpc.ID(),
	})
	if err != nil {
		return err
	}

	_, err = ec2.NewSecurityGroupRule(ctx, "db_ecs_allow_from_sg", &ec2.SecurityGroupRuleArgs{
		FromPort:              pulumi.Int(3306),
		Protocol:              pulumi.String("tcp"),
		SecurityGroupId:       sg.ID(),
		ToPort:                pulumi.Int(3306),
		Type:                  pulumi.String("ingress"),
		SourceSecurityGroupId: ecs.EcsSecurityGroup.ID(),
	})
	if err != nil {
		return err
	}

	_, err = ec2.NewSecurityGroupRule(ctx, "db_ecs_allow_to_all", &ec2.SecurityGroupRuleArgs{
		CidrBlocks: pulumi.StringArray{
			pulumi.String("0.0.0.0/0"),
		},
		FromPort:        pulumi.Int(0),
		Protocol:        pulumi.String("-1"),
		SecurityGroupId: sg.ID(),
		ToPort:          pulumi.Int(0),
		Type:            pulumi.String("egress"),
	})
	if err != nil {
		return err
	}

	subnetGroup, err := rds.NewSubnetGroup(ctx, "my_db_subnet_group", &rds.SubnetGroupArgs{
		Description: pulumi.String("description"),
		Name:        pulumi.String("sock-shop-mydbsubnetgroup-128kweik4u1y1"),
		SubnetIds: pulumi.StringArray{
			vpc.PublicSubnet1.ID(),
			vpc.PublicSubnet2.ID(),
		},
	})
	if err != nil {
		return err
	}

	_, err = rds.NewInstance(ctx, "catalogue", &rds.InstanceArgs{
		AutoMinorVersionUpgrade:    pulumi.Bool(true),
		CopyTagsToSnapshot:         pulumi.Bool(false),
		DeleteAutomatedBackups:     pulumi.Bool(true),
		Identifier:                 pulumi.String("sc1n069klknraxv"),
		InstanceClass:              pulumi.String("db.t2.medium"),
		MonitoringInterval:         pulumi.Int(0),
		PerformanceInsightsEnabled: pulumi.Bool(false),
		PubliclyAccessible:         pulumi.Bool(false),
		SkipFinalSnapshot:          pulumi.Bool(true),
		VpcSecurityGroupIds:        pulumi.StringArray{sg.ID()},
		DbSubnetGroupName:          subnetGroup.Name,
		AllocatedStorage:           pulumi.Int(100),
		Name:                       pulumi.String("socksdb"),
		Engine:                     pulumi.String("MySQL"),
		Username:                   pulumi.String("catalogue_user"),
		Password:                   pulumi.String("default_password"),
	})
	if err != nil {
		return err
	}

	return nil
}
